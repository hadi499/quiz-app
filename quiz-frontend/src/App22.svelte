<script>
  import { onMount } from "svelte";
  import StartScreen from "./lib/StartScreen.svelte";
  import QuestionCard from "./lib/QuestionCard.svelte";
  import ResultScreen from "./lib/ResultScreen.svelte";
  import AdminDashboard from "./lib/admin/AdminDashboard.svelte";
  import AuthForm from "./lib/AuthForm.svelte";
  import Modal from "./lib/Modal.svelte";

  let state = "start"; // 'start', 'quiz', 'result', 'admin', 'auth'
  let quizzes = [];
  let selectedQuiz = null;
  let questions = [];
  let currentIndex = 0;
  let score = 0;
  let userAnswers = []; // array of answers

  let modalConfig = {
    show: false,
    title: "",
    message: "",
    type: "alert",
    onConfirm: () => {},
  };
  function showAlert(msg) {
    modalConfig = {
      show: true,
      title: "Notification",
      message: msg,
      type: "alert",
      onConfirm: () => {},
    };
  }

  onMount(async () => {
    try {
      const res = await fetch("http://localhost:8080/api/quizzes");
      const data = await res.json();
      quizzes = data || [];
    } catch (err) {
      console.error(err);
      showAlert(
        "Failed to connect to backend server. Make sure it's running at http://localhost:8080.",
      );
    }
  });

  function selectQuiz(quiz) {
    if (!quiz.questions || quiz.questions.length === 0) {
      showAlert("This quiz has no questions yet!");
      return;
    }
    selectedQuiz = quiz;
    questions = quiz.questions;
    currentIndex = 0;
    score = 0;
    userAnswers = [];
    state = "quiz";
  }

  function handleAnswer(event) {
    const selected = event.detail.selected;
    const currentQ = questions[currentIndex];
    const isCorrect = selected === currentQ.answer;

    if (isCorrect) {
      score += 10;
    }

    userAnswers = [
      ...userAnswers,
      {
        question: currentQ,
        selected: selected,
        isCorrect: isCorrect,
        isTimeout: false,
      },
    ];

    nextQuestion();
  }

  function handleTimeout() {
    const currentQ = questions[currentIndex];
    userAnswers = [
      ...userAnswers,
      {
        question: currentQ,
        selected: -1,
        isCorrect: false,
        isTimeout: true,
      },
    ];
    nextQuestion();
  }

  function nextQuestion() {
    if (currentIndex + 1 < questions.length) {
      currentIndex++;
    } else {
      state = "result";
    }
  }

  function restartQuiz() {
    selectedQuiz = null;
    state = "start";
  }

  function goToAdmin() {
    const token = localStorage.getItem("token");
    if (token) {
      state = "admin";
    } else {
      state = "auth";
    }
  }

  function onLoginSuccess() {
    state = "admin";
  }

  function onBackFromAuth() {
    state = "start";
  }

  function exitAdmin() {
    state = "start";
    // refresh quizzes when exiting admin
    fetch("http://localhost:8080/api/quizzes")
      .then((res) => res.json())
      .then((data) => (quizzes = data || []))
      .catch((err) => console.error(err));
  }
</script>

<main>
  {#if state === "start"}
    <StartScreen {quizzes} onSelectQuiz={selectQuiz} onAdmin={goToAdmin} />
  {:else if state === "admin"}
    <AdminDashboard onExitAdmin={exitAdmin} />
  {:else if state === "auth"}
    <AuthForm {onLoginSuccess} onBack={onBackFromAuth} />
  {:else if state === "quiz"}
    {#if questions[currentIndex]}
      <QuestionCard
        question={questions[currentIndex]}
        on:answer={handleAnswer}
        on:timeout={handleTimeout}
      />
    {/if}
  {:else if state === "result"}
    <ResultScreen
      {score}
      total={questions.length * 10}
      {userAnswers}
      onRestart={restartQuiz}
    />
  {/if}
</main>

<Modal
  bind:show={modalConfig.show}
  title={modalConfig.title}
  message={modalConfig.message}
  type={modalConfig.type}
  onConfirm={modalConfig.onConfirm}
/>

<style>
  main {
    width: 100%;
  }
</style>
