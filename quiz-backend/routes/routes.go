package routes

import (
	"quiz-backend/controllers"
	"quiz-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// --- ROUTE API AUTH ---
	r.POST("/api/auth/register", controllers.Register)
	r.POST("/api/auth/login", controllers.Login)

	// --- ROUTE API QUIZZES ---
	r.GET("/api/quizzes", controllers.GetQuizzes)
	r.GET("/api/quizzes/:id", controllers.GetQuiz)

	// --- ROUTE API QUESTIONS ---
	r.GET("/api/questions", controllers.GetQuestions)
	r.GET("/api/questions/:id", controllers.GetQuestion)

	r.GET("/me", middleware.AuthMiddleware(), controllers.Me)

	// --- PROTECTED ROUTES ---
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/api/quizzes", controllers.CreateQuiz)
		protected.PUT("/api/quizzes/:id", controllers.UpdateQuiz)
		protected.DELETE("/api/quizzes/:id", controllers.DeleteQuiz)

		protected.POST("/api/questions", controllers.CreateQuestion)
		protected.PUT("/api/questions/:id", controllers.UpdateQuestion)
		protected.DELETE("/api/questions/:id", controllers.DeleteQuestion)
	}
}
