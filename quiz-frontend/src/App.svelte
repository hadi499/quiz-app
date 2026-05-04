<script>
  import { onMount } from "svelte";
  import StartScreen from "./lib/StartScreen.svelte";
  import QuestionCard from "./lib/QuestionCard.svelte";
  import ResultScreen from "./lib/ResultScreen.svelte";
  import AdminDashboard from "./lib/admin/AdminDashboard.svelte";
  import AuthForm from "./lib/AuthForm.svelte";
  import Modal from "./lib/Modal.svelte";

  const API_URL = import.meta.env.VITE_API_URL;

  let state = "start"; // 'start', 'quiz', 'result', 'admin', 'auth'
  let quizzes = [];
  let selectedQuiz = null;
  let questions = [];
  let currentIndex = 0;
  let score = 0;
  let userAnswers = []; // array of answers
  let user = null;

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
      const res = await fetch(`${API_URL}/api/quizzes`);
      const data = await res.json();
      quizzes = data || [];
    } catch (err) {
      console.error(err);
      showAlert(
        "Failed to connect to backend server. Make sure it's running at http://localhost:8080.",
      );
    }
  });

  async function getMe() {
    const token = localStorage.getItem("token");
    if (!token) return null;

    try {
      const res = await fetch(`${API_URL}/me`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) return null;

      return await res.json();
    } catch (err) {
      console.error(err);
      return null;
    }
  }
  onMount(async () => {
    const me = await getMe();

    if (!me) {
      localStorage.removeItem("token");
      user = null;
    } else {
      user = me;
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

  async function goToAdmin() {
    const user = await getMe();

    if (!user) {
      state = "auth";
      return;
    }

    if (user.role === "teacher") {
      state = "admin";
    } else {
      showAlert("Only teacher can access admin dashboard");
    }
  }

  async function onLoginSuccess() {
    user = await getMe();

    if (user?.role === "teacher") {
      state = "admin";
    } else {
      state = "start";
    }
  }

  function onBackFromAuth() {
    state = "start";
  }

  function handleLogout() {
    localStorage.removeItem("token");
    user = null; // ⬅️ penting
    state = "start";
  }

  function exitAdmin() {
    state = "start";
    // refresh quizzes when exiting admin
    fetch(`${API_URL}/api/quizzes`)
      .then((res) => res.json())
      .then((data) => (quizzes = data || []))
      .catch((err) => console.error(err));
  }
</script>

<main>
  {#if state === "start"}
    <StartScreen
      {quizzes}
      {user}
      onSelectQuiz={selectQuiz}
      onAdmin={goToAdmin}
      onLogout={handleLogout}
    />
  {:else if state === "admin"}
    <AdminDashboard onExitAdmin={exitAdmin} onLogout={handleLogout} />
  {:else if state === "auth"}
    <AuthForm {onLoginSuccess} onBack={onBackFromAuth} />
  {:else if state === "quiz"}
    {#if questions[currentIndex]}
      <QuestionCard
        question={questions[currentIndex]}
        timeLimit={selectedQuiz.timeLimit}
        questionNumber={currentIndex + 1}
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
