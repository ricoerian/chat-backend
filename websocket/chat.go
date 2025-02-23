package websocket

import (
	"chat-backend/config"
	"chat-backend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Struktur pesan yang dikirim ke WebSocket
type IncomingMessage struct {
	SenderID uint   `json:"sender_id"`
	Content  string `json:"content"`
}

// Struktur untuk mengirimkan pesan ke client
type OutgoingMessage struct {
	ID        uint   `json:"id"`
	SenderID  uint   `json:"sender_id"`
	Sender    string `json:"sender"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// WebSocket Server
type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	broadcast chan OutgoingMessage
	mutex     sync.Mutex
}

var wsServer = WebSocketServer{
	clients:   make(map[*websocket.Conn]bool),
	broadcast: make(chan OutgoingMessage),
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

// HandleConnections menangani koneksi WebSocket
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	wsServer.mutex.Lock()
	wsServer.clients[conn] = true
	wsServer.mutex.Unlock()

	log.Println("Client Connected")

	// Loop untuk membaca pesan dari WebSocket
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Client disconnected:", err)
			wsServer.mutex.Lock()
			delete(wsServer.clients, conn)
			wsServer.mutex.Unlock()
			break
		}

		// Parsing JSON dari pesan
		var incomingMessage IncomingMessage
		if err := json.Unmarshal(msg, &incomingMessage); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}

		// Cek apakah user dengan SenderID ada
		var user models.User
		if err := config.DB.First(&user, incomingMessage.SenderID).Error; err != nil {
			log.Println("User not found:", err)
			continue
		}

		// Simpan pesan ke database
		message := models.Message{
			SenderID: incomingMessage.SenderID,
			Content:  incomingMessage.Content,
		}
		if err := config.DB.Create(&message).Error; err != nil {
			log.Println("Failed to save message:", err)
			continue
		}

		fmt.Println("Received:", incomingMessage)

		// Konversi waktu ke zona WIB (GMT+7)
		loc, _ := time.LoadLocation("Asia/Jakarta")
		createdAtWIB := message.CreatedAt.In(loc).Format("2006-01-02 15:04:05")

		// Format pesan yang akan dikirim ke client
		outgoingMessage := OutgoingMessage{
			ID:        message.ID,
			SenderID:  message.SenderID,
			Sender:    user.Username,
			Content:   message.Content,
			CreatedAt: createdAtWIB,
		}

		// Broadcast pesan ke semua client
		wsServer.broadcast <- outgoingMessage
	}
}

// HandleMessages menangani pengiriman pesan ke semua client
func HandleMessages() {
	for {
		msg := <-wsServer.broadcast
		wsServer.mutex.Lock()
		for client := range wsServer.clients {
			jsonMessage, _ := json.Marshal(msg)
			err := client.WriteMessage(websocket.TextMessage, jsonMessage)
			if err != nil {
				log.Println("Error writing message:", err)
				client.Close()
				delete(wsServer.clients, client)
			}
		}
		wsServer.mutex.Unlock()
	}
};