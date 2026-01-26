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
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	port := os.Getenv("PORT")
	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        ":" + port,
		JWTsecret:   os.Getenv("JWT_SECRET"),
	}, nil
}
