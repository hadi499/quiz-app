<script>
  import { navigate } from "svelte-routing";
  import StartScreen from "../lib/StartScreen.svelte";
  import {
    appState,
    fetchQuizzes,
    startQuiz,
  } from "../stores/quiz.svelte.js";
  import { onMount } from "svelte";

  onMount(() => {
    if (appState.quizzes.length === 0) fetchQuizzes();
  });

  function selectQuiz(quiz) {
    if (startQuiz(quiz)) navigate("/quiz");
  }
</script>

<StartScreen quizzes={appState.quizzes} onSelectQuiz={selectQuiz} />
