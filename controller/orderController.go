package controller

import (
	"Depublic-App-Service/config"
	"Depublic-App-Service/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReserveTickets(c echo.Context) error {

	var requestBody struct {
		EventID      uint   `json:"event_id"`
		NumTickets   int    `json:"num_tickets"`
		CustomerName string `json:"customer_name"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Request")
	}

	user := c.Get("user").(map[string]interface{})
	customerName := user["name"].(string)

	var event model.Event
	if err := config.DB.First(&event, requestBody.EventID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Event Not found")
	}

	if event.AvailableTicket < requestBody.NumTickets {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ticket available")
	}

	order := model.Order{
		CustomerName: customerName,
	}
	if err := config.DB.Create(&order).Error; err != nil {
		return err
	}

	orderItems := make([]model.OrderItem, requestBody.NumTickets)
	for i := 0; i < requestBody.NumTickets; i++ {
		orderItems[i] = model.OrderItem{
			OrderID:   order.ID,
			EventID:   requestBody.EventID,
			EventName: event.EventName,
		}
	}
	if err := config.DB.Create(&orderItems).Error; err != nil {
		return err
	}

	event.AvailableTicket -= requestBody.NumTickets
	if err := config.DB.Save(&event).Error; err != nil {
		return err
	}

	var resnponseItems []map[string]interface{}
	for _, item := range orderItems {
		resnponseItems = append(resnponseItems, map[string]interface{}{
			"ticketId":  item.TicketID,
			"eventId":   item.EventID,
			"eventName": item.EventName,
			"orderId":   item.OrderID,
		})
	}

	response := map[string]interface{}{
		"message":    "Order successfully",
		"orderId":    order.ID,
		"orderItems": resnponseItems,
	}
	return c.JSON(http.StatusCreated, response)
}
