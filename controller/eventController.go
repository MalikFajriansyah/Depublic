package controller

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllEvent(c echo.Context) error {
	db := config.GetDB()
	var events []model.Event
	if err := db.Find(&events).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Event belum tersedia"})
	}

	return c.JSON(http.StatusOK, events)
}

func GetEventByCategory(c echo.Context) error {
	db := config.GetDB()

	category := c.Param("category")
	var events model.Event
	if err := db.Where("category = ?", category).Find(&events).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Kategori tersebut tidak ada di daftar"})
	}
	return c.JSON(http.StatusOK, events)
}

func GetEventByLocation(c echo.Context) error {
	db := config.GetDB()
	location := c.Param("location")
	var events model.Event
	if err := db.Where("location = ?", location).Find(&events).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Tidak ada untuk lokasi ini"})
	}
	return c.JSON(http.StatusOK, events)
}

func SearchEventName(c echo.Context) error {
	db := config.GetDB()

	searchQuery := c.QueryParam("event_name")
	var events []model.Event
	// if err := db.Where("event_name LIKE ?", "%"+searchQuery+"%").Find(&events); err != nil {
	// 	return c.JSON(http.StatusNotFound, map[string]string{"error": "Tidak ada event"})
	// }
	db.Where("event_name ILIKE ?", "%"+searchQuery+"%").Find(&events)
	return c.JSON(http.StatusOK, events)
}

func CreateEvent(c echo.Context) error {
	db := config.GetDB()

	events := new(model.Event)
	if err := c.Bind(events); err != nil {
		return err
	}

	db.Create(&events)

	return c.JSON(http.StatusCreated, map[string]string{"message": "Berhasil menambahkan event"})
}
