package controllers

import (
	"chat-backend/config"
	"chat-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoadMessages mengambil semua pesan dari database dan mengembalikannya sebagai JSON
func LoadMessages(c *gin.Context) {
	var messages []models.Message

	// Mengambil pesan dengan informasi pengirim (username)
	if err := config.DB.Preload("Sender").Order("created_at ASC").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil pesan"})
		return
	}

	// Konversi hasil agar `Sender.Username` bisa dikirim ke frontend
	var responseMessages []map[string]interface{}
	for _, msg := range messages {
		responseMessages = append(responseMessages, map[string]interface{}{
			"id":         msg.ID,
			"sender_id":  msg.SenderID,
			"sender":     msg.Sender.Username, // Ambil username
			"content":    msg.Content,
			"created_at": msg.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{"messages": responseMessages})
}
