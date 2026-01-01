package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close(context.Background())
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", conn)
			return next(c)
		}
	})
	e.GET("/student/:id", Getstudent)
	e.GET("/schedule/group/:id", GetGroupSchedule)
	e.GET("/all_class_schedule", GetAllSchedules)
	e.Logger.Fatal(e.Start(":8080"))
}

func Getstudent(c echo.Context) error {
	var name, surname, major, gender, groupName string
	var course_year int
	var birth_date time.Time
	id := c.Param("id")
	conc := c.Get("db").(*pgx.Conn)
	err := conc.QueryRow(c.Request().Context(), "SELECT s.firstname, s.lastname, g.groupname, s.major, s.course_year, s.gender, s.birth_date FROM students AS s JOIN groups as g ON s.group_id=g.id WHERE s.id=$1", id).Scan(&name, &surname, &groupName, &major, &course_year, &gender, &birth_date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	st := Student{
		Name:        name,
		Surname:     surname,
		GroupName:   groupName,
		Major:       major,
		Course_year: course_year,
		Gender:      gender,
		Birthd:      birth_date,
	}
	return c.JSON(http.StatusOK, st)
}

func GetGroupSchedule(c echo.Context) error {
	var groupName string
	var schedule []string
	id := c.Param("id")
	var Id int
	conc := c.Get("db").(*pgx.Conn)
	err := conc.QueryRow(c.Request().Context(), "SELECT g.id, g.groupname, c.lessons FROM groups AS g JOIN class_schedule AS c ON g.id=c.group_id WHERE g.id=$1", id).Scan(&Id, &groupName, &schedule)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	grSchedule := GroupSchedule{
		Id:        Id,
		GroupName: groupName,
		Lessons:   schedule,
	}
	return c.JSON(http.StatusOK, grSchedule)
}

func GetAllSchedules(c echo.Context) error {
	var schedule []GroupSchedule
	conc := c.Get("db").(*pgx.Conn)
	rows, err := conc.Query(c.Request().Context(), "SELECT g.id, g.groupname, c.lessons FROM groups as g JOIN class_schedule AS c ON g.id=c.group_id")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	defer rows.Close()
	for rows.Next() {
		var sch GroupSchedule
		if err := rows.Scan(&sch.Id, &sch.GroupName, &sch.Lessons); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": "Failed iteration"})
		}
		schedule = append(schedule, sch)
	}
	return c.JSON(http.StatusOK, schedule)
}
