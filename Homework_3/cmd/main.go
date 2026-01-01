package main

import (
	"context"
	"homework_3/internal/config"
	"homework_3/internal/handlers"
	"homework_3/internal/service"
	"homework_3/internal/storage"
	"homework_3/internal/storage/postgre"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

type Student struct {
	Name        string    `json:"firstname"`
	Surname     string    `json:"surname"`
	GroupName   string    `json:"group_name"`
	Major       string    `json:"major"`
	Course_year int       `json:"course_year"`
	Gender      string    `json:"gender"`
	Birthd      time.Time `json:"birth_date"`
}

type GroupSchedule struct {
	Id        int      `json:"id"`
	GroupName string   `json:"group_name"`
	Lessons   []string `json:"schedule"`
}

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
