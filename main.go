package main

import (
	"karmapay/routes"
	"log/slog"
	"os"
	"github.com/joho/godotenv"
	// "karma_auth/config"
)

func main() {
	err := godotenv.Load()
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))
	if err != nil {
		logger.Error("unable to load .env")
	}
	app := routes.Users()
	app.Listen(":9080")
}