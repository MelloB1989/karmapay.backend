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

type CreateCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	UID string `json:"uid"`
}

func CreateCustomer(c *fiber.Ctx) error {
	req := new(CreateCustomerRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to parse request body.",
			Data:    nil,
		})
	}
	cid, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890_-", 10);
	var customerData database.Customer = database.Customer{
		C_Email: req.Email,
		C_Name: req.FirstName+" "+req.LastName,
		C_Phone: req.Phone,
		CID: cid,
		C_Location: "NA",
		C_IP: "",
	}
	customers.CreateCustomer(customerData)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created user.",
		Data:    nil,
	})
}