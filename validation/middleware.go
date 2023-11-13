package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"

	"github.com/labstack/echo/v4"
)

func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	db := config.DB()
	var user model.User

	if db.Where("username = ? AND password = ?", username, password).First(&user) == nil {
		return false, nil
	}
	return true, nil
}
