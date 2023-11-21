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
	// e.Use(middleware.BasicAuth(validation.BasicAuthValidator))
	user := e.Group("/depublic")
	user.Use(middleware.BasicAuth(validation.BasicAuthValidator))
	user.POST("/login", controller.LoginUser)
	e.POST("/register", controller.RegisterUser)

	e.GET("/", controller.GetAllEvent)
	e.POST("/addEvent", controller.CreateEvent)
	e.Start(":8080")
}
