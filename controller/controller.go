package controller

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"Depublic-App-Service/validation"
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

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	// isValid, err := validation.BasicAuthValidator(u.Username, u.Password, c)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	// }
	// if !isValid {
	// 	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	// }

	var existingUser model.User
	if err := db.Where("username = ? AND password = ?", u.Username, u.Password).First(&existingUser).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}

	if !validation.VerifyPassword(u.Password, existingUser.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid username or password"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Login successfull"})
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
