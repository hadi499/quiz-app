<script>
  import { navigate } from "svelte-routing";
  import AuthForm from "../lib/AuthForm.svelte";
  import { appState, initAuth } from "../stores/quiz.svelte.js";

  async function onLoginSuccess() {
    await initAuth();
    if (appState.user?.role === "teacher") {
      navigate("/admin");
    } else {
      navigate("/");
    }
  }

  function onBack() {
    navigate("/");
  }
</script>

<AuthForm {onLoginSuccess} {onBack} />
