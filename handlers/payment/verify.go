package payment

import (
	// "fmt"
	// "net/http"
	"karmapay/database"
	"karmapay/helpers/orders"

	"github.com/gofiber/fiber/v2"
	rzutils "github.com/razorpay/razorpay-go/utils"
)


type VerifyPaymentRequest struct {
	OID string `json:"oid"`
	CID string `json:"cid"`
	PaymentID string `json:"payment_id"`
	Signature string `json:"signature"`
	RZKey string `json:"RZKey"`

}

func VerifyPayment(c *fiber.Ctx) error {
	req := new(VerifyPaymentRequest)
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
	params := map[string]interface{}{
		"razorpay_order_id": req.OID,
		"razorpay_payment_id": req.PaymentID,
	}
	
	signature := req.Signature;
	secret := req.RZKey;
	rzutils.VerifyPaymentSignature(params, signature, secret)
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