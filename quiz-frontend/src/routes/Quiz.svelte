<script>
  import { navigate } from "svelte-routing";
  import QuestionCard from "../lib/QuestionCard.svelte";
  import {
    appState,
    handleAnswer,
    handleTimeout,
    nextQuestion,
    advanceQuestion,
  } from "../stores/quiz.svelte.js";

  function onAnswer(e) {
    handleAnswer(e.detail.selected);
    if (nextQuestion()) {
      advanceQuestion();
    } else {
      navigate("/result");
    }
  }

  function onTimeout() {
    handleTimeout();
    if (nextQuestion()) {
      advanceQuestion();
    } else {
      navigate("/result");
    }
  }
</script>

{#if appState.questions[appState.currentIndex]}
  <QuestionCard
    question={appState.questions[appState.currentIndex]}
    timeLimit={appState.selectedQuiz.timeLimit}
    questionNumber={appState.currentIndex + 1}
    totalQuestions={appState.questions.length}
    on:answer={onAnswer}
    on:timeout={onTimeout}
  />
{:else}
  <div class="glass-card" style="text-align:center">
    <p>No questions loaded.</p>
    <a href="/" class="btn-primary" style="display:inline-block;text-decoration:none;">Back to Home</a>
  </div>
{/if}
