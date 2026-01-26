package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWTsecret   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        ":" + port,
		JWTsecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
