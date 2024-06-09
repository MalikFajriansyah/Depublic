package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerName string
	OrderItems   []OrderItem `gorm:"foreignkey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	EventID   uint
	EventName string
	TicketID  uint
}
