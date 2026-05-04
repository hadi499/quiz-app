<script>
  export let score = 0;
  export let total = 0;
  export let userAnswers = []; // array of { question, selected, isCorrect, isTimeout }
  export let onRestart;
</script>

<div class="glass-card result-card">
  <h1>Quiz Completed!</h1>
  <div class="score-display">
    <span class="score">{score}</span>
    <span class="total">/ {total}</span>
  </div>

  <div class="review-section">
    <h3>Review Answers</h3>
    {#each userAnswers as ans, i}
      <div class="review-item {ans.isCorrect ? 'correct' : 'wrong'}">
        <p class="q-text">
          {i + 1}.
          <span style="margin-left: 10px;color:#E6F082 ;"
            >{ans.question.question}</span
          >
        </p>
        <div class="details">
          {#if ans.isTimeout}
            <span class="badge error">Timeout</span>
          {:else}
            <span class="badge {ans.isCorrect ? 'success' : 'error'}">
              Your Answer: {ans.question.options[ans.selected]}
            </span>
          {/if}

          {#if !ans.isCorrect}
            <span class="badge success"
              >Correct: {ans.question.options[ans.question.answer]}</span
            >
          {/if}
        </div>
      </div>
    {/each}
  </div>

  <button class="btn-primary" style="margin-top: 2rem;" on:click={onRestart}
    >Restart Quiz</button
  >
</div>

<style>
  .result-card {
    max-height: 90vh;
    display: flex;
    flex-direction: column;
  }
  .score-display {
    text-align: center;
    margin: 1.5rem 0 2rem;
  }
  .score {
    font-size: 4rem;
    font-weight: 800;
    color: var(--accent-primary);
  }
  .total {
    font-size: 2rem;
    color: var(--text-muted);
  }
  h3 {
    margin-bottom: 1rem;
    color: var(--text-main);
  }
  .review-section {
    flex: 1;
    overflow-y: auto;
    padding-right: 0.5rem;
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  .review-item {
    background: rgba(0, 0, 0, 0.2);
    padding: 1rem;
    border-radius: 12px;
    border-left: 4px solid;
  }
  .review-item.correct {
    border-left-color: var(--success);
  }
  .review-item.wrong {
    border-left-color: var(--error);
  }
  .q-text {
    margin-bottom: 0.5rem;
    color: var(--text-main);
    font-weight: 500;
  }
  .details {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
  }
  .badge {
    padding: 0.25rem 0.75rem;
    border-radius: 999px;
    font-size: 0.875rem;
    font-weight: 500;
  }
  .badge.success {
    background: rgba(16, 185, 129, 0.2);
    color: #34d399;
  }
  .badge.error {
    background: rgba(239, 68, 68, 0.2);
    color: #f87171;
  }
</style>
