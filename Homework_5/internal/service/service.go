package service

import (
	"context"
	"errors"
	"homework_5/internal/models"
	"homework_5/internal/storage"
	"net/mail"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type StudentService struct {
	Storage   *storage.StudentStorage
	JWTsecret string
}

func New(st *storage.StudentStorage, js string) *StudentService {
	return &StudentService{Storage: st, JWTsecret: js}
}

func (s *StudentService) GetStudent(ctx context.Context, id string) (models.Student, error) {
	return s.Storage.GetStudent(ctx, id)
}

func (s *StudentService) GetGroupSchedule(ctx context.Context, id string) (models.GroupSchedule, error) {
	return s.Storage.GetGroupSchedule(ctx, id)
}

func (s *StudentService) GetAllSchedules(ctx context.Context) ([]models.GroupSchedule, error) {
	return s.Storage.GetAllSchedules(ctx)
}

func (s *StudentService) GetAttendanceStudent(ctx context.Context, id string) ([]models.Attendance, error) {
	return s.Storage.GetAttendanceStudent(ctx, id)
}

func (s *StudentService) GetAttendanceSubject(ctx context.Context, id string) ([]models.Attendance, error) {
	return s.Storage.GetAttendanceSubject(ctx, id)
}

func (s *StudentService) InsertAttendance(ctx context.Context, attendance models.Attendance) error {
	return s.Storage.InsertAttendance(ctx, attendance)
}

func (s *StudentService) RegisterUser(ctx context.Context, email, password string) error {
	if !checkEmail(email) {
		return errors.New("email is invalid")
	}
	pass_hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("invalid password")
	}
	return s.Storage.RegisterUser(ctx, email, string(pass_hash))
}

func (s *StudentService) LoginUser(ctx context.Context, email, password string) (string, error) {
	u, err := s.Storage.LoginUser(ctx, email)
	if err != nil {
		return "", errors.New("email is not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": u.ID,
		"expire":  time.Now().Add(time.Hour).Unix(),
	})
	return token.SignedString([]byte(s.JWTsecret))
}

func (s *StudentService) GetUserById(ctx context.Context, id int) (models.User, error) {
	return s.Storage.GetUserById(ctx, id)
}

func checkEmail(email string) bool {
	// Посмотрел в инете как лучше проверять почту и нашел стандартный пакет net/mail
	_, err := mail.ParseAddress(email)
	return err == nil
}
