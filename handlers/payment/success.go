package payment

import (
	// "fmt"
	// "net/http"
	"karmapay/database"
	"karmapay/helpers/orders"

	"github.com/gofiber/fiber/v2"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type SuccessPaymentRequest struct {
	OrderMode string `json:"mode"`
	OID string `json:"oid"`
	CID string `json:"cid"`
}

func SuccessPayment(c *fiber.Ctx) error {
	req := new(SuccessPaymentRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to parse request body.",
			Data:    nil,
		})
	}
	order, err := orders.GetOrderFromRedis(req.OID)
	if err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to get order from Redis.",
			Data:    nil,
		})
	}
	var orderData database.Order = database.Order{
		OrderID: order.OrderID,
		OrderAmount: order.OrderAmount,
		OrderCurrency: order.OrderCurrency,
		OrderDescription: order.OrderDescription,
		OrderStatus: "COMPLETED",
		OrderCID: req.CID,
		OrderTimeStamp: order.Timestamp,
		OrderUpiTransactionID: "",
		UID: order.UID,
	}
	orders.CreateOrder(orderData)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created user.",
		Data:    nil,
	})
}