<script>
  import { navigate } from "svelte-routing";
  import ResultScreen from "../lib/ResultScreen.svelte";
  import { appState, resetQuiz } from "../stores/quiz.svelte.js";

  function restart() {
    resetQuiz();
    navigate("/");
  }
</script>

<ResultScreen
  score={appState.score}
  total={appState.questions.length * 10}
  userAnswers={appState.userAnswers}
  onRestart={restart}
/>
