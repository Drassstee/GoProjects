package service

import (
	"context"
	"homework_3/internal/models"
	"homework_3/internal/storage"
)

type StudentService struct {
	Storage *storage.StudentStorage
}

func New(st *storage.StudentStorage) *StudentService {
	return &StudentService{Storage: st}
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
