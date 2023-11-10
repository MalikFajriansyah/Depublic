package middleware

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"

	"github.com/labstack/echo/v4"
)

func BasicAuth(roles ...string) echo.MiddlewareFunc {
	db := config.DB()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			username, password, ok := c.Request().BasicAuth()
			if ok {
				var user model.User
				if db.Where("username = ? AND password = ?", username, password).First(&user).Error == nil {
					for _, role := range roles {
						if user.Role == role {
							return next(c)
						}
					}
				}
			}
			return echo.ErrUnauthorized
		}
	}
}
