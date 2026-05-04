<script>
  import { onMount } from "svelte";
  import AdminQuestionForm from "./AdminQuestionForm.svelte";
  import Modal from "../Modal.svelte";

  export let selectedQuiz; // null for new
  export let onBack;

  let title = selectedQuiz ? selectedQuiz.title : "";
  let category = selectedQuiz ? selectedQuiz.category : "";
  let timeLimit = selectedQuiz ? selectedQuiz.timeLimit : "";

  let questions = [];
  let isEditingQuestion = false;
  let editingQuestion = null;

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

  async function fetchQuestions() {
    if (!selectedQuiz) return;
    const res = await fetch(
      `http://localhost:8080/api/quizzes/${selectedQuiz.id}`,
    );
    const data = await res.json();
    questions = (data.questions || []).sort((a, b) => b.id - a.id);
  }

  onMount(fetchQuestions);

  async function saveQuiz() {
    const payload = { title, category, timeLimit: parseInt(timeLimit) || 0 };
    const method = selectedQuiz ? "PUT" : "POST";
    const url = selectedQuiz
      ? `http://localhost:8080/api/quizzes/${selectedQuiz.id}`
      : `http://localhost:8080/api/quizzes`;

    const res = await fetch(url, {
      method,
      headers: getHeaders(),
      body: JSON.stringify({ title, category, timeLimit }),
    });

    if (res.ok) {
      if (!selectedQuiz) {
        const newQuiz = await res.json();
        selectedQuiz = newQuiz;
      } else {
        selectedQuiz.title = title;
        selectedQuiz.category = category;
        selectedQuiz.timeLimit = timeLimit;
      }
      showAlert("Quiz saved successfully!");
    }
  }

  function requestDeleteQuestion(id) {
    showConfirm("Delete this question?", async () => {
      const res = await fetch(`http://localhost:8080/api/questions/${id}`, {
        method: "DELETE",
        headers: getHeaders(),
      });
      if (res.ok) {
        fetchQuestions();
      } else {
        showAlert("Failed to delete question.");
      }
    });
  }

  function openQuestionForm(q = null) {
    editingQuestion = q;
    isEditingQuestion = true;
  }

  function closeQuestionForm(refresh = false) {
    isEditingQuestion = false;
    editingQuestion = null;
    if (refresh) fetchQuestions();
  }
</script>

{#if isEditingQuestion}
  <AdminQuestionForm
    question={editingQuestion}
    quizId={selectedQuiz.id}
    onClose={closeQuestionForm}
  />
{:else}
  <div class="header">
    <button class="btn-secondary small" on:click={onBack}>&larr; Back</button>
    <h2>{selectedQuiz ? "Edit Quiz" : "Create Quiz"}</h2>
  </div>

  <div class="form-group">
    <label for="title">Title</label>
    <input id="title" type="text" bind:value={title} class="form-control" />
  </div>

  <div class="form-group">
    <label for="category">Category</label>
    <input
      id="category"
      type="text"
      bind:value={category}
      class="form-control"
    />
  </div>
  <div class="form-group">
    <label for="timeLimit">Time Limit</label>
    <input
      id="timeLimit"
      type="number"
      bind:value={timeLimit}
      class="form-control"
    />
  </div>

  <button
    class="btn-primary small"
    on:click={saveQuiz}
    style="margin-bottom: 2rem;">Save Quiz Details</button
  >

  {#if selectedQuiz}
    <hr />
    <div class="header" style="margin-top: 2rem;">
      <h3>Questions</h3>
      <button class="btn-primary small" on:click={() => openQuestionForm(null)}
        >+ Add Question</button
      >
    </div>

    <div class="table-wrapper">
      <table class="data-table">
        <thead>
          <tr>
            <th>ID</th>
            <th>Question</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {#each questions as q}
            <tr>
              <td>{q.id}</td>
              <td>{q.question}</td>
              <td class="actions">
                <button
                  class="btn-primary small"
                  on:click={() => openQuestionForm(q)}>Edit</button
                >
                <button
                  class="btn-danger small"
                  on:click={() => requestDeleteQuestion(q.id)}>Delete</button
                >
              </td>
            </tr>
          {/each}
          {#if questions.length === 0}
            <tr>
              <td colspan="3" class="text-center">No questions added yet.</td>
            </tr>
          {/if}
        </tbody>
      </table>
    </div>
  {/if}
{/if}

<Modal
  bind:show={modalConfig.show}
  title={modalConfig.title}
  message={modalConfig.message}
  type={modalConfig.type}
  onConfirm={modalConfig.onConfirm}
/>

<style>
  .header {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1.5rem;
    justify-content: space-between;
  }
  .header h2,
  .header h3 {
    margin: 0;
  }
  .form-group {
    margin-bottom: 1.5rem;
  }
  label {
    display: block;
    margin-bottom: 0.5rem;
    color: var(--text-muted);
  }
  :global(.form-control) {
    width: 100%;
    padding: 0.75rem 1rem;
    background: rgba(0, 0, 0, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 8px;
    color: white;
    font-family: inherit;
    font-size: 1rem;
  }
  :global(.form-control:focus) {
    outline: none;
    border-color: var(--accent-primary);
  }
  hr {
    border: none;
    border-top: 1px solid rgba(255, 255, 255, 0.1);
  }
  .table-wrapper {
    overflow-x: auto;
  }
  .data-table {
    width: 100%;
    border-collapse: collapse;
  }
  .data-table td,
  .data-table th {
    padding: 0.75rem;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    text-align: left;
  }
  .actions {
    display: flex;
    gap: 0.5rem;
  }
  .text-center {
    text-align: center;
  }
</style>
