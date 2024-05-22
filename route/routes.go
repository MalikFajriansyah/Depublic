package route

import (
	"Depublic-App-Service/controller"
	"Depublic-App-Service/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRoutes() {
	e := echo.New()

	event := e.Group("/api")
	event.GET("/events", controller.GetAllEvent)
	event.GET("/events/category/:category", controller.GetEventByCategory)
	event.GET("/events/location/:location", controller.GetEventByLocation)
	event.GET("/events/search", controller.SearchEventName)
	event.POST("/addEvent", controller.CreateEvent)

	e.POST("/register", controller.RegisterUser)
	e.POST("/login", controller.LoginUser)

	restricted := e.Group("/restricted")
	restricted.Use(middlewares.Auth)
	restricted.GET("/home", controller.Home)
	e.Start(":8080")
}
