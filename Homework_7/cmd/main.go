package main

import (
	"context"
	"homework_6/internal/config"
	"homework_6/internal/handlers"
	"homework_6/internal/service"
	"homework_6/internal/storage"
	"homework_6/internal/storage/postgre"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := postgre.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())
	str := storage.New(db)
	srv := service.New(str, cfg.JWTsecret)
	hdl := handlers.New(srv, cfg.JWTsecret)
	server := echo.New()
	hdl.RegisterRoutes(server)
	server.Logger.Fatal(server.Start(cfg.Port))
}
