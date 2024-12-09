package handler

import (
	"go-chat-room/internal/db"
	"go-chat-room/internal/db/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(c *gin.Context) {
	// 1: Parse and validate input
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 2: Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error hashing password"})
		return
	}

	// 3: Save the user in the database
	user := model.User{
		Username: input.Username,
		Password: string(hashedPassword),
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user in database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": gin.H{"id": user.ID, "username": user.Username}})

}

// GET Users
func GetUsersHandler(c *gin.Context) {

	var users []model.User
	result := db.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users) // Return the actual data
}

// GET unique Users
func GetUsersByIdHandler(c *gin.Context) {
	var users []model.User
	id := c.Param("id")
	result := db.DB.First(&users, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
