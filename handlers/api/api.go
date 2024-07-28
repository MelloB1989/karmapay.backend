package api

import (
	// "fmt"
	// "net/http"
	"karmapay/helpers/api"
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

type CreateAPIRequest struct {
	APIKey string `json:"api_key"`
	PGEnum string `json:"pg_enum"`
}

func CreateAPI(c *fiber.Ctx) error {
	req := new(CreateAPIRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to parse request body.",
			Data:    nil,
		})
	}
	id, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890_-", 10);
	var API_KEY database.APIKeys = database.APIKeys{
		APIKey: req.APIKey,
		PGEnum: req.PGEnum,
		ID: id,
		UID: c.Locals("uid").(string),
	}
	api.CreateAPIKey(API_KEY)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created user.",
		Data:    nil,
	})
}