package main

import (
	"log"

	"quiz-backend/database"
	"quiz-backend/models"
	"quiz-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	// Koneksi ke Database
	database.Connect()

	// Migrasi otomatis untuk memastikan tabel ada
	database.DB.AutoMigrate(&models.User{}, &models.Quiz{}, &models.Question{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173","http://192.168.18.2:5173",}, // Alamat Svelte Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // <-- Ini kunci penyelesaiannya!
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Setup routes
	routes.SetupRoutes(r)

	log.Println("Server berjalan di port 8080...")
	r.Run(":8080")
}
