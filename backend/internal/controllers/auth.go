package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckHashedPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil // Which is true or false
}

func ValidateToken(token string) error {
	if token != "ACCESS_TOKEN" {
		return fmt.Errorf("token provided was invalid")
	}

	return nil
}

func AuthorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer")
	if err := ValidateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}