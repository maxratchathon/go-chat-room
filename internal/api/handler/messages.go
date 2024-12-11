package handler

import (
	"go-chat-room/internal/db"
	"go-chat-room/internal/db/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMessageHandler(c *gin.Context) {
	// 1: Parse and validate input
	var input struct {
		Content    string `json:"content" binding:"required"`
		SenderID   uint   `json:"senderId" binding:"required"`
		ChatRoomID uint   `json:"chatRoomId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 3: Save the user in the database
	message := model.Message{
		Content:    input.Content,
		SenderID:   input.SenderID,
		ChatRoomID: input.ChatRoomID,
	}
	if err := db.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message in database"})
		return
	}

	c.JSON(http.StatusOK, message)

}
