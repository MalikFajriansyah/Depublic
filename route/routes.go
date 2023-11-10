package route

import (
	"Depublic-App-Service/handler"
	"Depublic-App-Service/middleware"

	"github.com/labstack/echo/v4"
)

func initRoutes() {
	e := echo.New()

	userGroup := e.Group("/user")
	userGroup.Use(middleware.BasicAuth("user"))

	e.POST("/register", handler.RegisterUser)
	e.POST("/login", handler.LoginUser)

	e.Start(":8080")
}
