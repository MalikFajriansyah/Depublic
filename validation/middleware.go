package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	db := config.DatabaseInit()

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return false, nil
	}

	var existingUser model.User

	if err := db.Where("username = ?", username).First(&existingUser).Error; err != nil {
		message := (c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"}))
		return false, message
	}

	if err := verifyPasswordBasicAuth(password, existingUser.Password); err != nil {
		message := (c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"}))
		return false, message
	}
	return true, nil
}

// compare password dari user dengan passwod yang sudah terencrypt di database
func verifyPasswordBasicAuth(inputPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
