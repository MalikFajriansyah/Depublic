package controller

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

/* Func untuk user*/
func LoginUser(c echo.Context) error {
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
	config.DB.First(&user, "username = ? OR email = ?", body.Account, body.Account)

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to generate token",
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString

	cookie.SameSite = http.SameSiteLaxMode
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "user logged in",
	})
}

func RegisterUser(c echo.Context) error {
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

	result := config.DB.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed create user, email already exists",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "user created",
	})
}

func Home(c echo.Context) error {
	user := c.Get("user")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": user,
	})
}

/* Func untuk user*/
