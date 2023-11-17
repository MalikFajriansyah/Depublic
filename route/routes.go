package route

import (
	"Depublic-App-Service/controller"

	"github.com/labstack/echo/v4"
)

func InitRoutes() {
	e := echo.New()

	// e.Use(middleware.BasicAuth(validation.BasicAuthValidator))
	// e.Use(middleware.BasicAuth(validation.BasicAuthValidator))
	e.POST("/register", controller.RegisterUser)
	e.POST("/login", controller.LoginUser)
	e.Start(":8080")
}
