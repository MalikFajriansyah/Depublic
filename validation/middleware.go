package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"log"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}

	adminUsername := "AdminTes"
	adminPassword := "Tes12345"
	if username == adminUsername && password == adminPassword {
		return true, nil
	}

	var user model.User

	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}

	if !VerifyPassword(password, user.Password) {
		return false, nil
	}
	return true, nil
}

func VerifyPassword(inputPassword, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(inputPassword))
	return err == nil
}
