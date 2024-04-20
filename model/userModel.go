package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Role     string `json:"role" gorm:"default:user"`
}
