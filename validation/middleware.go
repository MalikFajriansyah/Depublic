package validation

import (
	"Depublic-App-Service/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	var user model.User

	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
