package controller

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

/* Func untuk user*/
func LoginUser(c echo.Context) error {
	db := config.DatabaseInit()

	var body struct {
		Account  string
		Password string
	}

	if c.Bind(&body) != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to read body",
		})
	}

	var user model.User
	db.First(&user, "username = ? OR email = ?", body.Account, body.Account)

	if user.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "user not found",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "password is incorrect",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user logged in",
	})
}

func RegisterUser(c echo.Context) error {
	db := config.DatabaseInit()

	var body struct {
		Username string
		Email    string
		Password string
	}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed generate password",
		})
	}

	user := model.User{Username: body.Username, Email: body.Email, Password: string(hash)}

	result := db.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed create user, email already exists",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "user created",
	})
}

/* Func untuk user*/
