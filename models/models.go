package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Messages []Message `gorm:"foreignKey:SenderID"`
}

type Message struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	SenderID  uint      `gorm:"not null" json:"sender_id"`
	Sender    User      `gorm:"foreignKey:SenderID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
