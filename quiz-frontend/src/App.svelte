<script>
  import { onMount } from "svelte";
  import { Router, Route, links, navigate } from "svelte-routing";
  import Modal from "./lib/Modal.svelte";
  import Home from "./routes/Home.svelte";
  import Quiz from "./routes/Quiz.svelte";
  import Result from "./routes/Result.svelte";
  import Login from "./routes/Login.svelte";
  import Admin from "./routes/Admin.svelte";
  import {
    appState,
    fetchQuizzes,
    initAuth,
    handleLogout,
  } from "./stores/quiz.svelte.js";

  onMount(async () => {
    fetchQuizzes();
    await initAuth();
  });

  function onLogout() {
    handleLogout();
    navigate("/");
  }

  export let url = "";
</script>

<div use:links>
  <Router {url}>
    <nav class="navbar">
      <div class="navbar-inner">
        <a href="/" class="nav-home">🏠 Home</a>
        <div class="nav-right">
          {#if !appState.user}
            <a href="/login" class="btn-secondary small"> 🔐 Login </a>
          {:else}
            <span class="nav-user">👋 {appState.user.username}</span>
            {#if appState.user.role === "teacher"}
              <a href="/admin" class="btn-secondary small"> ⚙️ Admin </a>
            {/if}
            <button class="btn-secondary small btn-logout" on:click={onLogout}>
              🚪 Logout
            </button>
          {/if}
        </div>
      </div>
    </nav>

    <main>
      <Route path="/">
        <Home />
      </Route>
      <Route path="/quiz">
        <Quiz />
      </Route>
      <Route path="/result">
        <Result />
      </Route>
      <Route path="/login">
        <Login />
      </Route>
      <Route path="/admin">
        <Admin />
      </Route>
      <Route path="*">
        <Home />
      </Route>
    </main>
  </Router>
</div>

<Modal
  bind:show={appState.modal.show}
  title={appState.modal.title}
  message={appState.modal.message}
  type={appState.modal.type}
  onConfirm={appState.modal.onConfirm}
/>

<style>
  main {
    width: 100%;
    padding-top: 70px;
  }
  .navbar {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 100;
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
    background: rgba(15, 23, 42, 0.92);
    backdrop-filter: blur(12px);
  }
  .navbar-inner {
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 0.75rem;
    padding: 0.75rem 1.5rem;
    max-width: 800px;
    margin: 0 auto;
  }
  .nav-home {
    color: var(--text-main);
    text-decoration: none;
    font-weight: 700;
    font-size: 1rem;
  }
  .nav-home:hover {
    color: var(--accent-primary);
  }
  .nav-right {
    display: flex;
    align-items: center;
    gap: 0.75rem;
  }
  .nav-user {
    font-weight: 600;
  }
  .btn-logout {
    color: #f87171;
    border-color: rgba(248, 113, 113, 0.3);
    text-decoration: none;
  }
</style>
