<script>
  import { fly, fade } from "svelte/transition";

  export let message = "";
  export let type = "success"; // "success" or "error"
  export let show = false;
  export let duration = 200;

  let timer;

  $: if (show) {
    clearTimeout(timer);
    timer = setTimeout(() => {
      show = false;
    }, duration);
  }
</script>

{#if show}
  <div
    class="flash-toast {type}"
    in:fly={{ y: -50, duration: 300 }}
    out:fade={{ duration: 300 }}
  >
    {message}
  </div>
{/if}

<style>
  .flash-toast {
    position: fixed;
    top: 24px;
    left: 50%;
    transform: translateX(-50%);
    padding: 0.75rem 2rem;
    border-radius: 999px;
    color: white;
    font-weight: 600;
    z-index: 2000;
    box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.5);
    backdrop-filter: blur(8px);
  }
  .flash-toast.success {
    background: rgba(16, 185, 129, 0.9);
    border: 1px solid #34d399;
  }
  .flash-toast.error {
    background: rgba(239, 68, 68, 0.9);
    border: 1px solid #f87171;
  }
</style>
