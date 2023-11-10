package handler

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	db := config.DB()

	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User
	if db.Where("username = ?", username).First(&user).RecordNotFound {
		return echo.NewHTTPError(http.StatusUnauthorized, "Username tidak dapat ditemukan")
	}

	if user.Password != password {
		return echo.NewHTTPError(http.StatusUnauthorized, "Kata sandi salah")
	}

	return c.JSON(http.StatusOK, user)
}

func RegisterUser(c echo.Context) error {
	db := config.DB()

	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return err
	}

	if newUser.Username == "" || newUser.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username dan password harus diisi")
	}

	var existingUser model.User
	if db.Where("username = ?", newUser.Username).First(&existingUser).Error == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Username sudah digunakan")
	}

	newUser.Role = "user"

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newUser)
}
