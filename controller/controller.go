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
	return c.JSON(http.StatusOK, map[string]string{"message": "Login successfull"})
}

func RegisterUser(c echo.Context) error {
	db := config.GetDB()

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

	//encrypt password
	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		return err
	}
	newUser.Password = hashedPassword

	//create data baru ke database
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Akun berhasil dibuat"})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

/* Func untuk user*/

/* Func untuk page JWT*/
func DashboardJwt(c echo.Context) error {

	var existingUser model.User

	if existingUser.Role == "admin" {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome Admin",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Login Success",
	})
}
