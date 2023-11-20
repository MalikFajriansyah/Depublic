package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
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

	// adminUsername := "AdminTes"
	// adminPassword := "Tes12345"
	// if username == adminUsername && password == adminPassword {
	// 	return true, nil
	// }

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return false, nil
	}

	var existingUser model.User

	if err := db.Where("username = ? AND password = ?", username, password).First(&existingUser).Error; err != nil {
		message := (c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"}))
		return false, message
	}
	return true, nil
}

// func VerifyPassword(inputPassword, storedHash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(inputPassword))
// 	return err == nil
// }
