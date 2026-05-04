<script>
  import Flash from "./Flash.svelte";

  export let onLoginSuccess; // function()
  export let onBack; // function()

  const API_URL = import.meta.env.VITE_API_URL;

  let mode = "login"; // 'login' or 'register'
  let username = "";
  let password = "";

  let flashConfig = { show: false, message: "", type: "error" };
  function showFlash(msg, type = "error") {
    flashConfig = { show: true, message: msg, type };
  }

  async function handleSubmit() {
    if (!username || !password) {
      showFlash("Please enter both username and password");
      return;
    }

    const endpoint =
      mode === "login" ? "/api/auth/login" : "/api/auth/register";

    try {
      const res = await fetch(`${API_URL}${endpoint}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();

      if (res.ok) {
        if (mode === "register") {
          showFlash("Registration successful! You can now log in.", "success");
          mode = "login";
          password = "";
        } else {
          // Login successful
          localStorage.setItem("token", data.token);
          onLoginSuccess();
        }
      } else {
        showFlash(data.error || "Authentication failed");
      }
    } catch (err) {
      showFlash("Failed to connect to server.");
    }
  }

  function toggleMode() {
    mode = mode === "login" ? "register" : "login";
    username = "";
    password = "";
  }
</script>

<div class="glass-card auth-card">
  <div class="header">
    <button class="btn-secondary small" on:click={onBack}>&larr; Back</button>
    <h2>{mode === "login" ? "Admin Login" : "Register Admin"}</h2>
  </div>

  <div class="form-group">
    <label for="username">Username</label>
    <input
      id="username"
      type="text"
      bind:value={username}
      class="form-control"
    />
  </div>

  <div class="form-group">
    <label for="password">Password</label>
    <input
      id="password"
      type="password"
      autocomplete="off"
      bind:value={password}
      class="form-control"
    />
  </div>

  <button class="btn-primary" on:click={handleSubmit} style="margin-top: 1rem;">
    {mode === "login" ? "Login" : "Register"}
  </button>

  <p class="toggle-text">
    {#if mode === "login"}
      Don't have an account? <span
        role="button"
        tabindex="0"
        on:click={toggleMode}
        on:keydown={(e) => e.key === "Enter" && toggleMode()}
        class="link">Register here</span
      >
    {:else}
      Already have an account? <span
        role="button"
        tabindex="0"
        on:click={toggleMode}
        on:keydown={(e) => e.key === "Enter" && toggleMode()}
        class="link">Login here</span
      >
    {/if}
  </p>
</div>

<Flash
  bind:show={flashConfig.show}
  message={flashConfig.message}
  type={flashConfig.type}
/>

<style>
  .auth-card {
    max-width: 400px;
    width: 100%;
    margin: 2rem auto;
  }
  .header {
    display: flex;
    align-items: center;
    gap: 1rem;
    margin-bottom: 2rem;
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
  .form-control {
    width: 100%;
    padding: 0.75rem 1rem;
    background: rgba(0, 0, 0, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 8px;
    color: white;
    font-family: inherit;
    font-size: 1rem;
  }
  .form-control:focus {
    outline: none;
    border-color: var(--accent-primary);
  }
  .toggle-text {
    margin-top: 1.5rem;
    text-align: center;
    font-size: 0.875rem;
  }
  .link {
    color: var(--accent-primary);
    cursor: pointer;
    font-weight: 600;
  }
  .link:hover {
    text-decoration: underline;
  }
</style>
