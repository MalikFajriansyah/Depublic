package validation

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func ProtectedEnpoint(c echo.Context) error {
	user := c.Get("user")
	fmt.Println(user)
	return c.String(http.StatusOK, "Welcome")
}

func LoginUseJwt(c echo.Context) error {
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

	//periksa data user
	var existingUser model.User
	if err := db.Where("username = ?", username).First(&existingUser).Error; err != nil {
		return echo.ErrUnauthorized
	}

	if err := verifyPasswordJwt(password, existingUser.Password); err != nil {
		return echo.ErrUnauthorized
	}
	claims := CustomClaims{
		existingUser.Username,
		jwt.StandardClaims{
			Id:        "user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte("SECRET"))
	if err != nil {
		return err
	}

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
	// 	StandardClaims: jwt.StandardClaims{},
	// 	Username:       existingUser.Username,
	// })

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func verifyPasswordJwt(inputPassword, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
}
