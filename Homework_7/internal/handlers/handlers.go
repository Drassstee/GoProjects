package handlers

import (
	"homework_6/internal/models"
	"homework_6/internal/service"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service   *service.StudentService
	JWTsecret string
}

func New(s *service.StudentService, js string) *Handler {
	return &Handler{Service: s, JWTsecret: js}
}

func (h *Handler) GetInstructor(c echo.Context) error {
	id := c.Param("id")
	st, err := h.Service.GetInstructor(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, st)
}

func (h *Handler) GetStudent(c echo.Context) error {
	id := c.Param("id")
	st, err := h.Service.GetStudent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, st)
}

func (h *Handler) GetStudentGrades(c echo.Context) error {
	id := c.Param("id")
	st, err := h.Service.GetStudentGrades(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, st)
}

func (h *Handler) GetGroupSchedule(c echo.Context) error {
	id := c.Param("id")
	sch, err := h.Service.GetGroupSchedule(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, sch)
}

func (h *Handler) GetAllSchedules(c echo.Context) error {
	schedules, err := h.Service.GetAllSchedules(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, schedules)
}

func (h *Handler) GetAttendanceStudent(c echo.Context) error {
	id := c.Param("id")
	at, err := h.Service.GetAttendanceStudent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, at)
}

func (h *Handler) GetAttendanceSubject(c echo.Context) error {
	id := c.Param("id")
	at, err := h.Service.GetAttendanceSubject(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, at)
}

func (h *Handler) InsertAttendance(c echo.Context) error {
	var at models.Attendance
	if err := c.Bind(&at); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}
	id, err := h.Service.InsertAttendance(c.Request().Context(), at)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"id": id})
}

func (h *Handler) RegisterUser(c echo.Context) error {
	var user models.AuthenticationReq
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}
	err := h.Service.RegisterUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, "success")
}

func (h *Handler) LoginUser(c echo.Context) error {
	var user models.AuthenticationReq
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}
	tok, err := h.Service.LoginUser(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": tok})
}

func (h *Handler) GetUserById(c echo.Context) error {
	id := c.Get("userId").(int)
	user, err := h.Service.GetUserById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *Handler) AuthMiddleware(n echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"err": "invalid token"})
		}
		str := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
			return []byte(h.JWTsecret), nil
		})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"err": "invalid token"})
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", int(claims["user_id"].(float64)))
		}
		return n(c)
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/student/:id", h.GetStudent)
	e.GET("/instructor/:id", h.GetInstructor)
	e.GET("/student/grades/:id", h.GetStudentGrades)
	e.GET("/schedule/group/:id", h.GetGroupSchedule)
	e.GET("/all_class_schedule", h.GetAllSchedules)
	e.GET("/attendanceBySubjectId/:id", h.GetAttendanceSubject)
	e.GET("/attendanceByStudentId/:id", h.GetAttendanceStudent)
	e.POST("/attendance/subject", h.InsertAttendance)
	e.POST("/api/auth/register", h.RegisterUser)
	e.POST("/api/auth/login", h.LoginUser)
	e.GET("/api/users/me", h.GetUserById, h.AuthMiddleware)
}
