package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaim struct {
	Name  string `json : "name"`
	Admin string `json : "admin"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	var username, password string
	db, err := config.DatabaseInit()
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	var existingUser model.User

	if err := db.Where("username = ? || password = ?", username, password).First(&existingUser).Error; err != nil {
		return echo.ErrUnauthorized
	}

	claims := &JwtCustomClaim{
		"admin",
		existingUser.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaim)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
