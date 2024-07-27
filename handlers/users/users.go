package handlers

import (
	// "fmt"
	// "net/http"
	"karmapay/helpers/users"
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

type SignupUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	BusinessName string `json:"business_name"`
	BusinessURL	string `json:"business_url"`
	Subdomain string `json:"subdomain"`
}

func CreateUser(c *fiber.Ctx) error {
	req := new(SignupUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to parse request body.",
			Data:    nil,
		})
	}
	uid, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890_-", 10);
	var userdata database.User = database.User{
		Username: req.Username,
		Password: req.Password,
		UID: uid,
		BusinessName: req.BusinessName,
		BusinessURL: req.BusinessURL,
		Subdomain: req.Subdomain,
	}
	users.CreateUser(userdata)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created user.",
		Data:    nil,
	})
}