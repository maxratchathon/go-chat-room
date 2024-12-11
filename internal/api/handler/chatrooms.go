package handler

import (
	"fmt"
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

// GET Chatroom
func GetChatRoomHandler(c *gin.Context) {

	var chatRoom []model.ChatRoom
	result := db.DB.Find(&chatRoom)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, chatRoom) // Return the actual data
}

// UPDATE Chatroom [name, createdBy]
func UpdateChatRoomHandler(c *gin.Context) {
	var input struct {
		Name      string `json:"name" binding:"required"`
		CreatedBy uint   `json:"createdBy" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	var chatRoom model.ChatRoom

	id := c.Param("id")
	db.DB.First(&chatRoom, id)

	chatRoom.Name = input.Name
	chatRoom.CreatedBy = input.CreatedBy

	result := db.DB.Save(&chatRoom)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": chatRoom.Name, "createdBy": chatRoom.CreatedBy})

}

// DELETE User by Id
func DeleteChatRoomHandler(c *gin.Context) {
	var chatRoom model.ChatRoom
	id := c.Param("id")
	db.DB.First(&chatRoom, id)
	result := db.DB.Delete(&chatRoom)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("the chatroom %s has been deleted successfully.", chatRoom.Name))
}
