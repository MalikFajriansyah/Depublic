package controller

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func GetAllEvent(c echo.Context) error {
	var events []model.Event
	if err := config.DB.Find(&events).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Event belum tersedia"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "All events",
		"data":    events,
	})
}

func CreateEvent(c echo.Context) error {
	var requestBody struct {
		EventName       string
		Date            time.Time
		Description     string
		Price           int
		Category        string
		Location        string
		AvailableTicket int
	}

	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	event := model.Event{
		EventName:       requestBody.EventName,
		Date:            requestBody.Date,
		Description:     requestBody.Description,
		Price:           requestBody.Price,
		Category:        requestBody.Category,
		Location:        requestBody.Location,
		AvailableTicket: requestBody.AvailableTicket,
	}

	result := config.DB.Create(&event)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed create new event",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "New event created successfully",
	})
}

func GetEventByCategory(c echo.Context) error {
	Category := c.Param("category")
	var events []model.Event
	if err := config.DB.Where("category = ?", Category).Find(&events).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Kategori tersebut tidak ada di daftar"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"category": Category,
		"data":     events,
	})
}

func GetEventByLocation(c echo.Context) error {
	location := c.Param("location")
	var events []model.Event
	if err := config.DB.Where("location = ?", location).Find(&events).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Tidak ada untuk lokasi ini"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "success",
		"location": location,
		"data":     events,
	})
}

func SearchEventName(c echo.Context) error {

	searchQuery := c.QueryParam("event_name")
	var events []model.Event
	// if err := db.Where("event_name LIKE ?", "%"+searchQuery+"%").Find(&events); err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]string{"error": "Tidak ada event"})
	// }
	config.DB.Where("event_name ILIKE ?", "%"+searchQuery+"%").Find(&events)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   events,
	})
}
