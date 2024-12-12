package model

// Message model for GORM
type Message struct {
	ID         uint   `gorm:"primaryKey"`
	Content    string `gorm:"type:text;not null"`
	SenderID   uint   `gorm:"not null"` // Foreign key to User
	ChatRoomID uint   `gorm:"not null"` // Foreign key to ChatRoom
}

// Users
type User struct {
	ID        uint       `gorm:"primaryKey"`
	Username  string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	ChatRooms []ChatRoom `gorm:"foreignKey:CreatedBy"`
	Messages  []Message  `gorm:"foreignKey:SenderID"`
}

type ChatRoom struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	CreatedBy uint      `gorm:"not null"` // Foreign key to User
	Messages  []Message `gorm:"foreignKey:ChatRoomID"`
}
