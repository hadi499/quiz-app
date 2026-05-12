<script>
  import { onMount } from "svelte";
  import AdminQuizForm from "./AdminQuizForm.svelte";
  import AdminQuestionForm from "./AdminQuestionForm.svelte";
  import Modal from "../Modal.svelte";

  export let onExitAdmin;
  export let onLogout;

  let quizzes = [];
  let allQuestions = [];
  let view = "list-quizzes"; // 'list-quizzes', 'list-questions', 'editQuiz', 'editQuestion'
  let selectedQuiz = null;
  let selectedQuestion = null;

  function getHeaders() {
    const token = localStorage.getItem("token");
    return {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    };
  }

  let modalConfig = {
    show: false,
    title: "",
    message: "",
    type: "alert",
    onConfirm: () => {},
  };
  function showAlert(msg) {
    modalConfig = {
      show: true,
      title: "Notification",
      message: msg,
      type: "alert",
      onConfirm: () => {},
    };
  }
  function showConfirm(msg, onYes) {
    modalConfig = {
      show: true,
      title: "Confirm Action",
      message: msg,
      type: "confirm",
      onConfirm: onYes,
    };
  }

  async function fetchQuizzes() {
    try {
      const res = await fetch("http://localhost:8080/api/quizzes");
      quizzes = (await res.json()) || [];
    } catch (err) {
      console.error(err);
    }
  }

  async function fetchAllQuestions() {
    try {
      const res = await fetch("http://localhost:8080/api/questions");
      const data = (await res.json()) || [];
      allQuestions = data.sort((a, b) => b.id - a.id);
    } catch (err) {
      console.error(err);
    }
  }

  onMount(() => {
    fetchQuizzes();
    fetchAllQuestions();
  });

  // -- Quiz Actions --
  function requestDeleteQuiz(id) {
    showConfirm("Are you sure you want to delete this quiz?", async () => {
      const res = await fetch(`http://localhost:8080/api/quizzes/${id}`, {
        method: "DELETE",
        headers: getHeaders(),
      });
      if (!res.ok) {
        if (res.status === 401) {
          showAlert("Unauthorized. Please log in again.");
          handleLogout();
        } else {
          showAlert("Failed to delete quiz.");
        }
        return;
      }
      fetchQuizzes();
      fetchAllQuestions();
    });
  }

  function createQuiz() {
    selectedQuiz = null;
    view = "editQuiz";
  }

  function editQuiz(quiz) {
    selectedQuiz = quiz;
    view = "editQuiz";
  }

  function onBackFromQuiz() {
    view = "list-quizzes";
    fetchQuizzes();
    fetchAllQuestions();
  }

  // -- Question Actions --
  function requestDeleteQuestion(id) {
    showConfirm("Are you sure you want to delete this question?", async () => {
      const res = await fetch(`http://localhost:8080/api/questions/${id}`, {
        method: "DELETE",
        headers: getHeaders(),
      });
      if (!res.ok) {
        if (res.status === 401) {
          showAlert("Unauthorized. Please log in again.");
          handleLogout();
        } else {
          showAlert("Failed to delete question.");
        }
        return;
      }
      fetchAllQuestions();
      fetchQuizzes();
    });
  }

  function createQuestion() {
    if (quizzes.length === 0) {
      showAlert("Please create a quiz first!");
      return;
    }
    selectedQuestion = null;
    view = "editQuestion";
  }

  function editQuestion(q) {
    selectedQuestion = q;
    view = "editQuestion";
  }

  function onBackFromQuestion(refresh) {
    view = "list-questions";
    if (refresh) {
      fetchAllQuestions();
      fetchQuizzes();
    }
  }

  function switchTab(tab) {
    view = tab;
    if (tab === "list-quizzes") fetchQuizzes();
    if (tab === "list-questions") fetchAllQuestions();
  }

  function handleLogout() {
    localStorage.removeItem("token");
    onExitAdmin();
  }
</script>

<div class="glass-card admin-container">
  {#if view === "list-quizzes" || view === "list-questions"}
    <div class="header">
      <h2>Admin Dashboard</h2>
      <div class="actions">
        <button class="btn-danger small" on:click={onLogout}>Logout</button>
        <button class="btn-secondary small" on:click={onExitAdmin}
          >Exit Admin</button
        >
      </div>
    </div>

    <div class="tabs">
      <button
        class="tab-btn {view === 'list-quizzes' ? 'active' : ''}"
        on:click={() => switchTab("list-quizzes")}>Quizzes</button
      >
      <button
        class="tab-btn {view === 'list-questions' ? 'active' : ''}"
        on:click={() => switchTab("list-questions")}>All Questions</button
      >
    </div>

    {#if view === "list-quizzes"}
      <div class="toolbar">
        <button class="btn-primary small" on:click={createQuiz}
          >+ New Quiz</button
        >
      </div>
      <div class="table-wrapper">
        <table class="data-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>QID</th>
              <th>Title</th>
              <th>Category</th>
              <th>Time</th>
              <th>Questions</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {#each quizzes as quiz, index}
              <tr>
                <td>{index + 1}</td>
                <td>{quiz.id}</td>
                <td>{quiz.title}</td>
                <td><span class="badge">{quiz.category}</span></td>
                <td>{quiz.timeLimit}</td>
                <td>{quiz.questions?.length || 0}</td>
                <td class="actions">
                  <button
                    class="btn-primary small"
                    on:click={() => editQuiz(quiz)}>Edit</button
                  >
                  <button
                    class="btn-danger small"
                    on:click={() => requestDeleteQuiz(quiz.id)}>Delete</button
                  >
                </td>
              </tr>
            {/each}
            {#if quizzes.length === 0}
              <tr>
                <td colspan="5" class="text-center"
                  >No quizzes found. Create one!</td
                >
              </tr>
            {/if}
          </tbody>
        </table>
      </div>
    {:else if view === "list-questions"}
      <div class="toolbar">
        <button class="btn-primary small" on:click={createQuestion}
          >+ New Question</button
        >
      </div>
      <div class="table-wrapper">
        <table class="data-table">
          <thead>
            <tr>
              <th>ID</th>
              <th>Quiz ID</th>
              <th>DB Id</th>
              <th>Question Text</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {#each allQuestions as q, index}
              <tr>
                <td>{index + 1}</td>
                <td><span class="badge">{q.quiz_id}</span></td>
                <td><span class="badge">{q.id}</span></td>
                <td class="truncate">{q.question}</td>
                <td class="actions">
                  <button
                    class="btn-primary small"
                    on:click={() => editQuestion(q)}>Edit</button
                  >
                  <button
                    class="btn-danger small"
                    on:click={() => requestDeleteQuestion(q.id)}>Delete</button
                  >
                </td>
              </tr>
            {/each}
            {#if allQuestions.length === 0}
              <tr>
                <td colspan="4" class="text-center">No questions found.</td>
              </tr>
            {/if}
          </tbody>
        </table>
      </div>
    {/if}
  {:else if view === "editQuiz"}
    <AdminQuizForm {selectedQuiz} onBack={onBackFromQuiz} />
  {:else if view === "editQuestion"}
    <AdminQuestionForm
      question={selectedQuestion}
      {quizzes}
      quizId={selectedQuestion ? selectedQuestion.quiz_id : quizzes[0].id}
      onClose={onBackFromQuestion}
    />
  {/if}
</div>

<Modal
  bind:show={modalConfig.show}
  title={modalConfig.title}
  message={modalConfig.message}
  type={modalConfig.type}
  onConfirm={modalConfig.onConfirm}
/>

<style>
  .admin-container {
    max-width: 900px;
    width: 100%;
    max-height: 90vh;
    overflow-y: auto;
  }
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.5rem;
  }
  .tabs {
    display: flex;
    gap: 1rem;
    margin-bottom: 1.5rem;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    padding-bottom: 0.5rem;
  }
  .tab-btn {
    background: none;
    border: none;
    color: var(--text-muted);
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    padding: 0.5rem 1rem;
    transition: color 0.2s;
  }
  .tab-btn.active {
    color: var(--accent-primary);
    border-bottom: 2px solid var(--accent-primary);
  }
  .tab-btn:hover {
    color: white;
  }
  .toolbar {
    margin-bottom: 1rem;
    display: flex;
    justify-content: flex-end;
  }
  .table-wrapper {
    overflow-x: auto;
  }
  .data-table {
    width: 100%;
    border-collapse: collapse;
  }
  .data-table th,
  .data-table td {
    padding: 1rem;
    text-align: left;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }
  .data-table th {
    font-weight: 600;
    color: var(--text-muted);
  }
  .truncate {
    max-width: 300px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }
  .actions {
    display: flex;
    gap: 0.5rem;
  }
  .badge {
    background: rgba(59, 130, 246, 0.2);
    color: #60a5fa;
    padding: 0.25rem 0.5rem;
    border-radius: 4px;
    font-size: 0.8rem;
  }
  .text-center {
    text-align: center;
  }
</style>
