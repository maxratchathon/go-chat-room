package main

import (
	"go-chat-room/internal/api/handler"
	"go-chat-room/internal/controllers"
	"go-chat-room/internal/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Websocket Upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all origins for simplicity
	},
}

func wsHandler(c *gin.Context) {
	// Upgrade HTTP connection to websocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket: ", err)
		return
	}
	defer conn.Close()

	log.Println("client connected")

	for {
		// Read message from client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received: %s\n", message)

		// Echo message back to client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	// Initialize Gin router
	router := gin.Default()

	protected := router.Group("/", controllers.AuthorizationMiddleware)

	// Define a api routes
	router.GET("/ws", wsHandler)

	// login api
	router.POST("/login", handler.LoginHandler)

	// users CRUD api
	protected.GET("/users", handler.GetUsersHandler)
	protected.GET("/users/:id", handler.GetUsersByIdHandler)
	protected.POST("/users", handler.CreateUserHandler)
	protected.PATCH("/users/:id", handler.UpdateUserByIdHandler)
	protected.DELETE("/users/:id", handler.DeleteUserByIdHandler)

	// chatroom CRUD api
	protected.POST("/chat-rooms", handler.CreateChatRoomHandler)
	protected.GET("/chat-rooms", handler.GetChatRoomHandler)
	protected.PATCH("/chat-rooms/:id", handler.UpdateChatRoomHandler)
	protected.DELETE("/chat-rooms/:id", handler.DeleteChatRoomHandler)

	// Message CRUD api
	protected.POST("/messages", handler.CreateMessageHandler)
	protected.GET("/messages", handler.GetMessageHandler)

	// Init DB
	dsn := "host=localhost user=postgres password=secret dbname=go-chat-rooms port=5432 sslmode=disable"

	_, err := db.InitDB(dsn)
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	// defer db.DB.Close()

	// Migrate DB schemas
	db.MigrateDB()

	// Start the server
	port := ":8080"
	log.Println("Sever running on port", port)
	if err := router.Run(port); err != nil {
		log.Fatal("Server failed:", err)
	}
}
