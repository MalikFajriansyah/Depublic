package model

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	EventID uint `json:"event_id"`
}
