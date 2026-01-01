package main

import (
	"context"
	"homework_3/internal/config"
	"homework_3/internal/handlers"
	"homework_3/internal/service"
	"homework_3/internal/storage"
	"homework_3/internal/storage/postgre"
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
	srv := service.New(str)
	hdl := handlers.New(srv)
	server := echo.New()
	hdl.RegisterRoutes(server)
	server.Logger.Fatal(server.Start(cfg.Port))
}
