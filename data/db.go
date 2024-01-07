package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, _ = gorm.Open(sqlite.Open("resource/chat.sqlite"))

func InitDB() {
	DB.AutoMigrate(User{})
	DB.AutoMigrate(Chat{})
	DB.AutoMigrate(Message{})
}
