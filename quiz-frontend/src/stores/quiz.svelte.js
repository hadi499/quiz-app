const API_URL = import.meta.env.VITE_API_URL;

export const appState = $state({
  quizzes: [],
  selectedQuiz: null,
  questions: [],
  currentIndex: 0,
  score: 0,
  userAnswers: [],
  user: null,
  modal: {
    show: false,
    title: "",
    message: "",
    type: "alert",
    onConfirm: () => {},
  },
});

export function showAlert(msg) {
  appState.modal = {
    show: true,
    title: "Notification",
    message: msg,
    type: "alert",
    onConfirm: () => {},
  };
}

export async function fetchQuizzes() {
  try {
    const res = await fetch(`${API_URL}/api/quizzes`);
    appState.quizzes = (await res.json()) || [];
  } catch (err) {
    console.error(err);
    showAlert("Failed to connect to backend server. Make sure it's running.");
  }
}

export async function getMe() {
  const token = localStorage.getItem("token");
  if (!token) return null;

  try {
    const res = await fetch(`${API_URL}/me`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    if (!res.ok) return null;
    return await res.json();
  } catch (err) {
    console.error(err);
    return null;
  }
}

export async function initAuth() {
  const me = await getMe();
  if (!me) {
    localStorage.removeItem("token");
    appState.user = null;
  } else {
    appState.user = me;
  }
}

export function handleLogout() {
  localStorage.removeItem("token");
  appState.user = null;
}

export function startQuiz(quiz) {
  if (!quiz.questions || quiz.questions.length === 0) {
    showAlert("This quiz has no questions yet!");
    return false;
  }
  appState.selectedQuiz = quiz;
  appState.questions = quiz.questions;
  appState.currentIndex = 0;
  appState.score = 0;
  appState.userAnswers = [];
  return true;
}

export function handleAnswer(selected) {
  const currentQ = appState.questions[appState.currentIndex];
  const isCorrect = selected === currentQ.answer;

  if (isCorrect) appState.score += 10;

  appState.userAnswers = [
    ...appState.userAnswers,
    { question: currentQ, selected, isCorrect, isTimeout: false },
  ];
}

export function handleTimeout() {
  const currentQ = appState.questions[appState.currentIndex];
  appState.userAnswers = [
    ...appState.userAnswers,
    { question: currentQ, selected: -1, isCorrect: false, isTimeout: true },
  ];
}

export function nextQuestion() {
  return appState.currentIndex + 1 < appState.questions.length;
}

export function advanceQuestion() {
  appState.currentIndex++;
}

export function resetQuiz() {
  appState.selectedQuiz = null;
  appState.questions = [];
  appState.currentIndex = 0;
  appState.score = 0;
  appState.userAnswers = [];
}
