package mysql

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	RoomID    []Room `gorm:"many2many:user_rooms"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Room struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"unique"`
}

type Chat struct {
	ID      int64
	From    string
	To      string
	Message string
}
