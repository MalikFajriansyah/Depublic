package route

import (
	"Depublic-App-Service/controller"
	"Depublic-App-Service/validation"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes() {
	e := echo.New()

	event := e.Group("/api")
	event.GET("/events", controller.GetAllEvent)
	event.GET("/events/category/:category", controller.GetEventByCategory)
	event.GET("/events/location/:location", controller.GetEventByLocation)
	event.GET("/events/search", controller.SearchEventName)
	event.POST("/addEvent", controller.CreateEvent)
	// e.Use(echojwt.JWT([]byte("SECRET")))

	// Basic Authentication
	basicAuth := e.Group("/basicAuth")
	basicAuth.Use(middleware.BasicAuth(validation.BasicAuthValidator))
	basicAuth.POST("/login", controller.LoginUser)

	// JWT Authentication
	jwt := e.Group("/jwt")
	jwt.Use(echojwt.WithConfig(echojwt.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte("SECRET"),
	}))
	// role user
	jwt.GET("/home", controller.DashboardJwt)

	e.POST("/register", controller.RegisterUser)
	e.POST("/login", validation.LoginUseJwt)
	e.Start(":8080")
}
