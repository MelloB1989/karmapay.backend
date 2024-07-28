package handlers

import (
	// "fmt"
	// "net/http"
	"karmapay/config"
	"karmapay/database"
	"karmapay/helpers/users"
	"log"

	"github.com/dgrijalva/jwt-go"
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

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

func LoginUser(c *fiber.Ctx) error {
	req := new(LoginUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to parse request body.",
			Data:    nil,
		})
	}
	user := users.GetUserByUsername(req.Username)
	if user.Password == req.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": user.Username,
			"uid": user.UID,
		})
		// Sign and get the complete encoded token as a string using the secret
		jwtSecret := []byte(config.NewConfig().JWTSecret)
        tokenString, err := token.SignedString(jwtSecret)
        if err != nil {
            log.Println("Failed to create token:", err)
            return c.JSON(ResponseHTTP{
                Success: false,
                Message: "Failed to create token.",
                Data:    nil,
            })
        }
		return c.JSON(ResponseHTTP{
			Success: true,
			Message: "Successfully logged in.",
			Data:    tokenString,
		})
	} else {
		return c.JSON(ResponseHTTP{
			Success: false,
			Message: "Invalid credentials.",
			Data:    nil,
		})
	}
}