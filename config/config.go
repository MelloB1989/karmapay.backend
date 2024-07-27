package config

import (
	"os"
	"github.com/joho/godotenv"
	"log/slog"
)

type Config struct {
	Port string
	AdminKey string
	JWTSecret string
}

func NewConfig() *Config {
	err := godotenv.Load()
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))
	if err != nil {
		logger.Error("unable to load .env")
	}
	return &Config{
		Port: os.Getenv("PORT"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		AdminKey: os.Getenv("ADMIN_KEY"),
	}
}