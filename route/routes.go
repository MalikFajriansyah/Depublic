package route

import (
	"Depublic-App-Service/controller"
	"Depublic-App-Service/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes() {
	e := echo.New()

	// e.Use(middleware.BasicAuth(validation.BasicAuthValidator))
	e.POST("/register", controller.RegisterUser)
	e.POST("/login", controller.LoginUser, middleware.BasicAuth(validation.BasicAuthValidator))
	e.Start(":8080")
}
