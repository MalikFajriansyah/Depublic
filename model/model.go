package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:user"`
}

type Ticket struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
}

type Event struct {
	gorm.Model
	Name      string  `json:"name" gorm:"not null"`
	Date      string  `json:"date" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`
	Available int     `json:"available" gorm:"not null"`
}
