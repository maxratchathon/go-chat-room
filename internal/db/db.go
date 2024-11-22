package db

import (
	"fmt"
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

// Message model for GORM
type Message struct {
	ID      uint   `gorm:"primaryKey"`
	Content string `gorm:"type:text"`
	UserID  uint   `gorm:"not null"`
	RoomID  uint   `gorm:"not null"`
}

func MigrateDB() {
	// Automatically migrate the schema (create tables)
	err := DB.AutoMigrate(&Message{})
	if err != nil {
		log.Fatalf("Falied to migrate database: %v", err)

	}
	log.Println("Database migration complete")
}
