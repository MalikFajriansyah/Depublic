package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func BasicAuthValidator(username, password string, c echo.Context) (bool, error) {
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	adminUsername := "AdminTes"
	adminPassword := "Tes12345"
	if username == adminUsername && password == adminPassword {
		message := c.JSON(http.StatusOK, map[string]string{"message": "Welcome to dashboard admin"})
		return true, message
	}

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return false, nil
	}

	var existingUser model.User

	if err := db.Where("username = ?", username).First(&existingUser).Error; err != nil {
		message := (c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"}))
		return false, message
	}

	if err := verifyPassword(password, existingUser.Password); err != nil {
		message := (c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"}))
		return false, message
	}
	return true, nil
}

// compare password dari user dengan passwod yang sudah terencrypt di database
func verifyPassword(inputPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
