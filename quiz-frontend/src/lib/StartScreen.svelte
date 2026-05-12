<script>
  export let quizzes = [];
  export let onSelectQuiz; // function(quiz)
</script>

<div class="header-section text-center">
  <h1>Les Balonggarut</h1>
  <h2>Select a Quiz</h2>
  <p>Choose a category and test your knowledge!</p>
</div>

{#if quizzes.length === 0}
  <div class="glass-card text-center">
    <p>Loading quizzes... (make sure backend is running)</p>
  </div>
{:else}
  <div class="quiz-grid">
    {#each quizzes as quiz}
      <div
        class="glass-card quiz-card"
        role="button"
        tabindex="0"
        on:click={() => onSelectQuiz(quiz)}
        on:keydown={(e) => {
          if (e.key === "Enter" || e.key === " ") onSelectQuiz(quiz);
        }}
      >
        <span class="category-badge">{quiz.category}</span>
        <h2>{quiz.title}</h2>
        <p>{quiz.questions?.length || 0} Questions</p>
        <button tabindex="-1" class="btn-primary" style="margin-top: 1rem;"
          >Start Quiz</button
        >
      </div>
    {/each}
  </div>
{/if}

<style>
  .text-center {
    text-align: center;
  }
  .header-section {
    margin-bottom: 2rem;
  }
  .quiz-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 1.5rem;
    width: 100%;
  }
  .quiz-card {
    cursor: pointer;
    transition:
      transform 0.2s ease,
      box-shadow 0.2s ease;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
  }
  .quiz-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 15px 30px -10px rgba(59, 130, 246, 0.4);
    border-color: var(--accent-primary);
  }
  .quiz-card h2 {
    margin-top: 1rem;
    margin-bottom: 0.5rem;
    font-size: 1.25rem;
  }
  .category-badge {
    background: rgba(59, 130, 246, 0.2);
    color: #60a5fa;
    padding: 0.25rem 0.75rem;
    border-radius: 999px;
    font-size: 0.875rem;
    font-weight: 600;
  }
</style>
