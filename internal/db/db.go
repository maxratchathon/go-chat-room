package db

import (
	"fmt"
	"go-chat-room/internal/db/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB instance
var DB *gorm.DB

// Initialize the database connection
func InitDB(dsn string) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database %w", err)
	}
	log.Println("Database connection success")
	return DB, nil
}

func MigrateDB() {
	// Automatically migrate the schema (create tables)
	err := DB.AutoMigrate(
		&model.User{},
		&model.ChatRoom{},
		&model.Message{},
	)
	if err != nil {
		log.Fatalf("Falied to migrate database: %v", err)

	}
	log.Println("Database migration complete")
}
