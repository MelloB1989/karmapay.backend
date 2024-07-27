package handlers

import (
	// "fmt"
	// "net/http"
	"karmapay/helpers/customers"
	"karmapay/database"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type CreateOrderRequest struct {
	OrderAmount string `json:"order_amt"`
	OrderCurrency string `json:"order_currency"`
	OrderDescription string `json:"order_description"`
	OrderMode string `json:"order_mode"`
	WebhookURL string `json:"webhook_url"`
	RedirectURL string `json:"redirect_url"`
	Registration string `json:"registration"`
}

func CreateOrder(c *fiber.Ctx) error {
	req := new(CreateOrderRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to parse request body.",
			Data:    nil,
		})
	}
	oid, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890_-", 10);
	var orderData database.RedisOrder = database.RedisOrder{
		OrderID: oid,
		OrderStatus: "PENDING",
	}
	customers.CreateCustomer(customerData)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created user.",
		Data:    nil,
	})
}