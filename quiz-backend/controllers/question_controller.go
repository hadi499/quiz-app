package controllers

import (
	"net/http"
	"quiz-backend/database"
	"quiz-backend/models"

	"github.com/gin-gonic/gin"
)

// GET ALL: Mengambil semua pertanyaan
func GetQuestions(c *gin.Context) {
	var questions []models.Question

	if err := database.DB.Order("id asc").Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, questions)
}

// GET ONE: Mengambil satu pertanyaan berdasarkan ID
func GetQuestion(c *gin.Context) {
	var question models.Question
	id := c.Param("id")
	if err := database.DB.First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pertanyaan tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, question)
}

// CREATE: Menambah pertanyaan baru
func CreateQuestion(c *gin.Context) {
	var newQuestion models.Question
	if err := c.ShouldBindJSON(&newQuestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek jumlah pertanyaan untuk quiz ini
	var count int64
	database.DB.Model(&models.Question{}).Where("quiz_id = ?", newQuestion.QuizID).Count(&count)
	if count >= 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Kuis ini sudah mencapai batas maksimal 10 pertanyaan"})
		return
	}

	if err := database.DB.Create(&newQuestion).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newQuestion)
}

// UPDATE: Mengubah pertanyaan yang sudah ada
func UpdateQuestion(c *gin.Context) {
	var question models.Question
	id := c.Param("id")

	// Cek apakah data ada
	if err := database.DB.First(&question, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pertanyaan tidak ditemukan"})
		return
	}

	// Bind data baru dari JSON ke struct terpisah
	var input models.Question
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pertahankan ID asli dari URL
	input.ID = question.ID

	if err := database.DB.Save(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, input)
}

// DELETE: Menghapus pertanyaan
func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Question{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pertanyaan berhasil dihapus"})
}
