package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	EventName       string    `json:"event_name" gorm:"not null"`
	Date            time.Time `json:"date" gorm:"not null"` //format time CCYY-MM-DDThh:mm:ssZ Z = zona UTC
	Description     string    `json:"description" gorm:"not null"`
	Price           int       `json:"price" gorm:"not null"`
	Category        string    `json:"category" gorm:"not null"`
	Location        string    `json:"location" gorm:"not null"`
	AvailableTicket int       `json:"available_ticket" gorm:"not null"`
}
