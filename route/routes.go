package route

import (
	"Depublic-App-Service/controller"
	"Depublic-App-Service/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitRoutes() {
	e := echo.New()

	userGroup := e.Group("/user")
	userGroup.Use(middleware.BasicAuth("user"))

	e.POST("/register", controller.RegisterUser)
	e.POST("/login", controller.LoginUser)

	e.Start(":8080")
}
