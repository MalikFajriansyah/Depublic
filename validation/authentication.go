package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func LoginUseJwt(c echo.Context) error {

	db := config.DatabaseInit()

	identifier := c.FormValue("username or email")
	password := c.FormValue("password")

	//periksa data user
	var existingUser model.User

	if err := db.Where("username = ? OR email = ?", identifier, identifier).First(&existingUser).Error; err != nil {
		return echo.ErrUnauthorized
	}

	if err := verifyPasswordJwt(password, existingUser.Password); err != nil {
		return echo.ErrUnauthorized
	}

	claims := CustomClaims{
		existingUser.Username,
		existingUser.Role,
		jwt.StandardClaims{
			Id:        "user_id",
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte("SECRET"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Berhasil",
		"token":   token,
	})
}

func verifyPasswordJwt(inputPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
