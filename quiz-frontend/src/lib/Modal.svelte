<script>
  export let show = false;
  export let title = "Notification";
  export let message = "";
  export let type = "alert"; // "alert" or "confirm"
  export let onConfirm = () => {};
  export let onCancel = () => {};

  function handleConfirm() {
    show = false;
    onConfirm();
  }

  function handleCancel() {
    show = false;
    onCancel();
  }
</script>

{#if show}
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <div class="modal-backdrop" on:click={handleCancel}>
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="modal-content glass-card" on:click|stopPropagation>
      <h3>{title}</h3>
      <p>{message}</p>
      <div class="modal-actions">
        {#if type === "confirm"}
          <button class="btn-secondary small" on:click={handleCancel}
            >Cancel</button
          >
          <button class="btn-danger small" on:click={handleConfirm}
            >Confirm</button
          >
        {:else}
          <button class="btn-primary small" on:click={handleConfirm}>OK</button>
        {/if}
      </div>
    </div>
  </div>
{/if}

<style>
  .modal-backdrop {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(4px);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
  }
  .modal-content {
    min-width: 300px;
    max-width: 400px;
    padding: 1.5rem;
    animation: scaleIn 0.2s ease-out;
  }
  .modal-content h3 {
    margin-top: 0;
    margin-bottom: 1rem;
    color: var(--text-main);
  }
  .modal-content p {
    margin-bottom: 1.5rem;
    color: var(--text-muted);
  }
  .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 0.5rem;
  }
  @keyframes scaleIn {
    from {
      transform: scale(0.95);
      opacity: 0;
    }
    to {
      transform: scale(1);
      opacity: 1;
    }
  }
</style>
