<script>
  import { navigate } from "svelte-routing";
  import AdminDashboard from "../lib/admin/AdminDashboard.svelte";
  import { appState, fetchQuizzes, handleLogout } from "../stores/quiz.svelte.js";

  async function exitAdmin() {
    await fetchQuizzes();
    navigate("/");
  }

  function onLogout() {
    handleLogout();
    navigate("/");
  }
</script>

<AdminDashboard onExitAdmin={exitAdmin} {onLogout} />
