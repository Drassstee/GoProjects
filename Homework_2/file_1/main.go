package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/postreq", check)
	e.Logger.Fatal(e.Start(":1323"))
}

func check(c echo.Context) error {
	per := new(Person)
	if err := c.Bind(per); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}
	if per.Name == "" || per.Surname == "" {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Name and surname cannot be empty"})
	}
	return c.JSON(http.StatusOK, per)
}
