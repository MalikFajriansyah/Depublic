package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"log"

	"github.com/labstack/echo/v4"
)

func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}

	var user model.User

	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
