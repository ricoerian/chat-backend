package main

import (
	"chat-backend/config"
	"chat-backend/controllers"
	"chat-backend/websocket"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://chat-frontend-khaki.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Endpoint untuk registrasi dan login
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/messages", controllers.LoadMessages)

	r.GET("/ws", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Izinkan akses dari semua origin
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	
		websocket.HandleConnections(c.Writer, c.Request)
	})

	// Menjalankan WebSocket listener
	go websocket.HandleMessages()

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
