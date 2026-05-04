<script>
  import { createEventDispatcher, onMount, onDestroy } from "svelte";
  const dispatch = createEventDispatcher();

  export let question = { id: 0, question: "", options: [] };
  export let timeLimit = 10;

  // 1. Tambahkan prop baru untuk nomor urut (default 1)
  export let questionNumber = 1;

  let timeLeft;
  let timer;

  function startTimer() {
    clearInterval(timer);
    timeLeft = timeLimit;
    timer = setInterval(() => {
      timeLeft -= 1;
      if (timeLeft <= 0) {
        clearInterval(timer);
        dispatch("timeout");
      }
    }, 1000);
  }

  $: question, startTimer();

  onDestroy(() => {
    clearInterval(timer);
  });

  function selectOption(idx) {
    clearInterval(timer);
    dispatch("answer", { selected: idx });
  }

  // Sekalian saya perbaiki ulang rumus progress bar-nya
  $: progressWidth = (timeLeft / timeLimit) * 100;
</script>

<div class="glass-card">
  <h3 style="margin-bottom: 15px; text-align: center;color: #FFC81E;">
    No. {questionNumber}
  </h3>
  <div class="timer-bar">
    <div class="timer-fill" style="width: {progressWidth}%"></div>
  </div>
  <div class="header">
    <span class="time-text">{timeLeft}s</span>
  </div>

  <h2>{question.question}</h2>

  <div class="options">
    {#each question.options as option, idx}
      <button class="option-btn" on:click={() => selectOption(idx)}>
        {option}
      </button>
    {/each}
  </div>
</div>

<style>
  .timer-bar {
    width: 100%;
    height: 6px;
    background: rgba(255, 255, 255, 0.1);
    border-radius: 3px;
    margin-bottom: 1.5rem;
    overflow: hidden;
  }
  .timer-fill {
    height: 100%;
    background: var(--accent-primary);
    transition: width 1s linear;
  }
  .header {
    display: flex;
    justify-content: flex-end;
    margin-bottom: 1rem;
  }
  .time-text {
    font-weight: 700;
    color: var(--accent-primary);
  }
  .options {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
    margin-top: 2rem;
  }
  .option-btn {
    padding: 1rem 1.5rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 12px;
    color: var(--text-main);
    font-size: 1rem;
    text-align: left;
    cursor: pointer;
    transition: all 0.2s ease;
    font-family: inherit;
  }
  .option-btn:hover {
    background: rgba(255, 255, 255, 0.1);
    border-color: var(--accent-primary);
    transform: translateX(4px);
  }
</style>
