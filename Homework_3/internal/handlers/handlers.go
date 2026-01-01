package handlers

import (
	"homework_3/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *service.StudentService
}

func New(s *service.StudentService) *Handler {
	return &Handler{Service: s}
}

func (s *Handler) GetStudent(c echo.Context) error {
	id := c.Param("id")
	st, err := s.Service.GetStudent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, st)
}

func (s *Handler) GetGroupSchedule(c echo.Context) error {
	id := c.Param("id")
	sch, err := s.Service.GetGroupSchedule(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, sch)
}

func (s *Handler) GetAllSchedules(c echo.Context) error {
	schedules, err := s.Service.GetAllSchedules(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, schedules)
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/student/:id", h.GetStudent)
	e.GET("/schedule/group/:id", h.GetGroupSchedule)
	e.GET("/all_class_schedule", h.GetAllSchedules)
}
