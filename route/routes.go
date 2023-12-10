package route

import (
	"Depublic-App-Service/controller"
	"Depublic-App-Service/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes() {
	e := echo.New()

	user := e.Group("/depublic")
	user.Use(middleware.BasicAuth(validation.BasicAuthValidator))
	user.POST("/login", controller.LoginUser)
	e.POST("/register", controller.RegisterUser)

	e.GET("/events", controller.GetAllEvent)
	e.GET("/events/category/:category", controller.GetEventByCategory)
	e.GET("/events/location/:location", controller.GetEventByLocation)
	e.GET("/events/search", controller.SearchEventName)
	e.POST("/addEvent", controller.CreateEvent)
	e.Start(":8080")
}
