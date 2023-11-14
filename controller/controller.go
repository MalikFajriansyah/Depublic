package controller

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User
	if db.Where("username = ?", username).First(&user) == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Username tidak ditemukan")
	}

	if user.Password != password {
		return echo.NewHTTPError(http.StatusUnauthorized, "Kata sandi salah")
	}

	return c.JSON(http.StatusOK, user)
}

func RegisterUser(c echo.Context) error {
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	var newUser model.User
	if err := c.Bind(&newUser); err != nil {
		return err
	}

	//validasi data
	if newUser.Username == "" || newUser.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "username dan password harus diisi")
	}

	//cek jika data user sudah terdaftar
	var existingUser model.User
	if db.Where("username = ?", newUser.Username).First(&existingUser).Error == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "username sudah digunakan")
	}

	//setel role user registrasi
	newUser.Role = "user"

	//create data baru ke database
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newUser)
}
