package models

import (
	"time"

	"gorm.io/gorm"
)

// Gunakan nama unik agar tidak bentrok dengan model lain
type ChatMessage struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Sender    string    `json:"sender"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
