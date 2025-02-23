package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Messages []Message `gorm:"foreignKey:SenderID"`
}

// Message model
type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	SenderID  uint      `gorm:"not null" json:"sender_id"`
	Sender    User      `gorm:"foreignKey:SenderID"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// BeforeCreate: Set zona waktu sebelum menyimpan ke database
func (m *Message) BeforeCreate(tx *gorm.DB) (err error) {
	loc, err := time.LoadLocation("Asia/Jakarta") // Pastikan zona waktu benar
	if err != nil {
		return err
	}
	m.CreatedAt = time.Now().In(loc) // Set waktu sesuai zona waktu Indonesia
	return nil
}
