package handler

import (
	"go-chat-room/internal/db"
	"go-chat-room/internal/db/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateChatRoomHandler(c *gin.Context) {
	// 1: Parse and validate input
	var input struct {
		Name   string `json:"name" binding:"required"`
		UserId uint   `json:"userId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 3: Save the user in the database
	chatRoom := model.ChatRoom{
		Name:      input.Name,
		CreatedBy: input.UserId,
	}
	if err := db.DB.Create(&chatRoom).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create chatroom in database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "chatroom": gin.H{"id": chatRoom.ID, "name": chatRoom.Name, "createdBy": chatRoom.CreatedBy}})

}

// GET Users
func GetChatRoomHandler(c *gin.Context) {

	var chatRoom []model.ChatRoom
	result := db.DB.Find(&chatRoom)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, chatRoom) // Return the actual data
}
