package main

import (
	"context"
	"homework_7/internal/config"
	"homework_7/internal/handlers"
	"homework_7/internal/migrations"
	"homework_7/internal/service"
	"homework_7/internal/storage"
	"homework_7/internal/storage/postgre"
	"log"

	_ "homework_7/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	eSwag "github.com/swaggo/echo-swagger"
)

// @title Deployment of university app and swagger just for info
// @description Swagger API
// @host university-management-app-oojg.onrender.com
// @BasePath /
// @Schemes https
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
	if err = migrations.Run(context.Background(), db); err != nil {
		log.Fatal(err)
	}
	str := storage.New(db)
	srv := service.New(str, cfg.JWTsecret)
	hdl := handlers.New(srv, cfg.JWTsecret)
	server := echo.New()
	server.Use(middleware.CORS())
	server.GET("/swagger/*", eSwag.WrapHandler)
	hdl.RegisterRoutes(server)
	server.Logger.Fatal(server.Start(cfg.Port))
}
