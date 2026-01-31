package handlers

import (
	"homework_7/internal/models"
	"homework_7/internal/service"
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

// GetInstructor godoc
// @Summary      Get instructor info
// @Description  Get instructor info by ID
// @Tags         instructor
// @Produce      json
// @Param id path string true "Instructor Id"
// @Success      200  {object}  models.Instructor
// @Failure      500  {object}  map[string]any
// @Router       /instructor/{id} [get]
func (h *Handler) GetInstructor(c echo.Context) error {
	id := c.Param("id")
	st, err := h.Service.GetInstructor(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, st)
}

// GetStudent godoc
// @Summary      Get student info
// @Description  Get student info by ID
// @Tags         student
// @Produce      json
// @Param id path string true "Student Id"
// @Success      200  {object}  models.Student
// @Failure      500  {object}  map[string]any
// @Router       /student/{id} [get]
func (h *Handler) GetStudent(c echo.Context) error {
	id := c.Param("id")
	st, err := h.Service.GetStudent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, st)
}

// InsertStudent godoc
// @Summary      Insert Student Info
// @Description  Insert Student Info
// @Tags         student
// @Accept       json
// @Produce      json
// @Param input body models.StudentReq true "Student Info"
// @Success      200  {object}  map[string]any
// @Failure      400  {object}  map[string]any
// @Failure      500  {object}  map[string]any
// @Router       /student/create [post]
func (h *Handler) InsertStudent(c echo.Context) error {
	var st models.StudentReq
	if err := c.Bind(&st); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}
	id, err := h.Service.InsertStudent(c.Request().Context(), st)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"id": id})
}

// InsertGroup godoc
// @Summary      Insert Group Info
// @Description  Insert Group Info
// @Tags         group
// @Accept       json
// @Produce      json
// @Param input body models.GroupReq true "Group Info"
// @Success      200  {object}  map[string]any
// @Failure      400  {object}  map[string]any
// @Failure      500  {object}  map[string]any
// @Router       /group/create [post]
func (h *Handler) InsertGroup(c echo.Context) error {
	var gr models.GroupReq
	if err := c.Bind(&gr); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}
	id, err := h.Service.InsertGroup(c.Request().Context(), gr.Groupname)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"id": id})
}

// GetStudentGrades godoc
// @Summary      Get student's grades info
// @Description  Get student's grades by ID
// @Tags         grades
// @Produce      json
// @Param id path string true "Student Id"
// @Success      200  {object}  []models.Grade
// @Failure      500  {object}  map[string]any
// @Router       /student/grades/{id} [get]
func (h *Handler) GetStudentGrades(c echo.Context) error {
	id := c.Param("id")
	st, err := h.Service.GetStudentGrades(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, st)
}

// GetGroupSchedule godoc
// @Summary      Get group's schedule
// @Description  Get group's schedule by ID
// @Tags         schedule
// @Produce      json
// @Param id path string true "Group Id"
// @Success      200  {object}  models.GroupSchedule
// @Failure      500  {object}  map[string]any
// @Router       /schedule/group/{id} [get]
func (h *Handler) GetGroupSchedule(c echo.Context) error {
	id := c.Param("id")
	sch, err := h.Service.GetGroupSchedule(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, sch)
}

// GetAllSchedules godoc
// @Summary      Get all groups' schedules
// @Description  Get all groups' schedules
// @Tags         schedule
// @Produce      json
// @Success      200  {object}  []models.GroupSchedule
// @Failure      500  {object}  map[string]any
// @Router       /all_class_schedule [get]
func (h *Handler) GetAllSchedules(c echo.Context) error {
	schedules, err := h.Service.GetAllSchedules(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err2: ": err.Error()})
	}
	return c.JSON(http.StatusOK, schedules)
}

// GetAttendanceStudent godoc
// @Summary      Get student's attendance info
// @Description  Get student's attendance info by ID
// @Tags         attendace
// @Produce      json
// @Param id path string true "Student Id"
// @Success      200  {object}  []models.Attendance
// @Failure      500  {object}  map[string]any
// @Router       /attendanceByStudentId/{id} [get]
func (h *Handler) GetAttendanceStudent(c echo.Context) error {
	id := c.Param("id")
	at, err := h.Service.GetAttendanceStudent(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, at)
}

// GetAttendanceSubject godoc
// @Summary      Get attendance info of the subject
// @Description  Get attendance info of the subject by ID
// @Tags         attendace
// @Produce      json
// @Param id path string true "Subject Id"
// @Success      200  {object}  []models.Attendance
// @Failure      500  {object}  map[string]any
// @Router       /attendanceBySubjectId/{id} [get]
func (h *Handler) GetAttendanceSubject(c echo.Context) error {
	id := c.Param("id")
	at, err := h.Service.GetAttendanceSubject(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, at)
}

// InsertAttendance godoc
// @Summary      Insert attendance info
// @Description  Insert attendance info
// @Tags         attendace
// @Accept       json
// @Produce      json
// @Param input body models.Attendance true "Attendance Info"
// @Success      200  {object}  map[string]any
// @Failure      400  {object}  map[string]any
// @Failure      500  {object}  map[string]any
// @Router       /attendance/subject [post]
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

// RegisterUser godoc
// @Summary      Register user
// @Description  Handler for Registration
// @Tags         auth
// @Produce      json
// @Param input body models.AuthenticationReq true "Account Credentials"
// @Success      200  {object} map[string]string
// @Failure      400  {object} map[string]any
// @Failure      401  {object} map[string]any
// @Router       /api/auth/register [post]
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

// LoginUser godoc
// @Summary      Login user
// @Description  Handler for Login
// @Tags         auth
// @Produce      json
// @Param input body models.LoginReq true "Account Credentials"
// @Success      200  {object} map[string]string
// @Failure      400  {object} map[string]any
// @Failure      401  {object} map[string]any
// @Router       /api/auth/login [post]
func (h *Handler) LoginUser(c echo.Context) error {
	var user models.LoginReq
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": err.Error()})
	}
	tok, err := h.Service.LoginUser(c.Request().Context(), user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{"err: ": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"token": tok})
}

// GetUserById godoc
// @Summary      Get current user info
// @Security     BearerAuth
// @Tags         users
// @Produce      json
// @Success      200  {object}  models.User
// @Router       /api/users/me [get]
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
	e.POST("/student/create", h.InsertStudent)
	e.POST("/group/create", h.InsertGroup)
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
