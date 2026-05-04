<script>
  import Flash from "../Flash.svelte";

  export let question = null; // null for new
  export let quizId = null;
  export let quizzes = null; // Optional: array of quizzes for dropdown
  export let onClose; // function(shouldRefresh)

  let qText = question ? question.question : "";
  let options =
    question && question.options ? [...question.options] : ["", "", "", ""];
  while (options.length < 4) options.push("");

  let answer = question ? question.answer : 0;
  let selectedQuizId = question
    ? question.quiz_id
    : quizId || (quizzes && quizzes.length > 0 ? quizzes[0].id : null);

  function getHeaders() {
    const token = localStorage.getItem("token");
    return {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    };
  }

  let flashConfig = { show: false, message: "", type: "error" };
  function showFlash(msg, type = "error") {
    flashConfig = { show: true, message: msg, type: type };
  }

  async function saveQuestion() {
    if (!qText || options.some((o) => !o)) {
      showFlash("Please fill in the question and all 4 options.", "error");
      return;
    }

    if (!selectedQuizId) {
      showFlash("Please select a quiz for this question.", "error");
      return;
    }

    const payload = {
      quiz_id: parseInt(selectedQuizId),
      question: qText,
      options: options,
      answer: parseInt(answer),
    };

    const method = question ? "PUT" : "POST";
    const url = question
      ? `http://localhost:8080/api/questions/${question.id}`
      : `http://localhost:8080/api/questions`;

    const res = await fetch(url, {
      method,
      headers: getHeaders(),
      body: JSON.stringify(payload),
    });

    if (res.ok) {
      showFlash("Question saved successfully!", "success");
      setTimeout(() => onClose(true), 700); // Wait briefly before closing
    } else {
      const err = await res.json();
      showFlash(`Failed: ${err.error || "Unknown error"}`, "error");
    }
  }
</script>

<div class="header">
  <button class="btn-secondary small" on:click={() => onClose(false)}
    >&larr; Back</button
  >
  <h2>{question ? "Edit Question" : "Add Question"}</h2>
</div>

{#if quizzes && quizzes.length > 0}
  <div class="form-group">
    <label for="quizSelect">Assign to Quiz</label>
    <select id="quizSelect" bind:value={selectedQuizId} class="form-control">
      {#each quizzes as q}
        <option value={q.id}>{q.title}</option>
      {/each}
    </select>
  </div>
{/if}

<div class="form-group">
  <label for="qText">Question Text</label>
  <textarea id="qText" bind:value={qText} class="form-control" rows="3"
  ></textarea>
</div>

<h4>Options</h4>
{#each options as opt, i}
  <div class="form-group option-row">
    <input
      type="radio"
      name="correctAnswer"
      value={i}
      bind:group={answer}
      title="Mark as correct answer"
    />
    <input
      type="text"
      bind:value={options[i]}
      class="form-control"
      placeholder={`Option ${i + 1}`}
    />
  </div>
{/each}
<p class="help-text">Select the radio button next to the correct answer.</p>

<button
  class="btn-primary small"
  on:click={saveQuestion}
  style="margin-top: 1rem;">Save Question</button
>

<Flash
  bind:show={flashConfig.show}
  message={flashConfig.message}
  type={flashConfig.type}
/>

<style>
  .header {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 1.5rem;
    justify-content: flex-start;
  }
  .header h2 {
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
  h4 {
    margin-bottom: 1rem;
  }
  .option-row {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 0.5rem;
  }
  .option-row input[type="radio"] {
    transform: scale(1.5);
    accent-color: var(--success);
    cursor: pointer;
  }
  .help-text {
    font-size: 0.875rem;
    color: var(--text-muted);
    margin-bottom: 1rem;
  }
  select.form-control {
    cursor: pointer;
    appearance: auto;
  }
  select.form-control option {
    background: var(--bg-card);
    color: white;
  }
</style>
