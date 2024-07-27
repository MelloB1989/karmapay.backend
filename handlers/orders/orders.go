package handlers

import (
	// "fmt"
	// "net/http"
	api "karmapay/helpers/api"
	users "karmapay/helpers/users"
	orders "karmapay/helpers/orders"
	"karmapay/database"
	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"encoding/json"
	"fmt"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type CreateOrderRequest struct {
	OrderAmount float64 `json:"order_amt"`
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
		UID: c.Locals("uid").(string),
		Email: c.Locals("email").(string),
		API_KEY: api.GetAPIKeyByUIDAndPGEnum(c.Locals("uid").(string), req.OrderMode).APIKey,
		OrderAmount: fmt.Sprintf("%.9f", req.OrderAmount),
		OrderCurrency: req.OrderCurrency,
		OrderDescription: req.OrderDescription,
		Subdomain: users.GetUserByUID(c.Locals("uid").(string)).Subdomain,
		OrderMode: req.OrderMode,
		OrderCID: "",
		PGOrder: json.RawMessage(`{}`),
		KPAPI: c.Locals("kpapi").(string),
		Registration: req.Registration,
		RedirectURL: req.RedirectURL,
	}
	orders.PushOrderToRedis(orderData)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created order.",
		Data:    json.RawMessage(`{"oid": "` + oid + `"}`),
	})
}