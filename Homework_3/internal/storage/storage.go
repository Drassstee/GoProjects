package storage

import (
	"context"
	"homework_3/internal/models"

	"github.com/jackc/pgx/v5"
)

type StudentStorage struct {
	DB *pgx.Conn
}

func New(db *pgx.Conn) *StudentStorage {
	return &StudentStorage{DB: db}
}

func (s *StudentStorage) GetStudent(ctx context.Context, id string) (models.Student, error) {
	var st models.Student
	err := s.DB.QueryRow(ctx, "SELECT s.firstname, s.lastname, g.groupname, s.major, s.course_year, s.gender, s.birth_date FROM students AS s JOIN groups as g ON s.group_id=g.id WHERE s.id=$1", id).Scan(&st.Name, &st.Surname, &st.GroupName, &st.Major, &st.Course_year, &st.Gender, &st.Birthd)
	return st, err
}

func (s *StudentStorage) GetGroupSchedule(ctx context.Context, id string) (models.GroupSchedule, error) {
	var sch models.GroupSchedule
	err := s.DB.QueryRow(ctx, "SELECT g.id, g.groupname, c.lessons FROM groups AS g JOIN class_schedule AS c ON g.id=c.group_id WHERE g.id=$1", id).Scan(&sch.Id, &sch.GroupName, &sch.Lessons)
	return sch, err
}

func (s *StudentStorage) GetAllSchedules(ctx context.Context) ([]models.GroupSchedule, error) {
	var schedules []models.GroupSchedule
	rows, err := s.DB.Query(ctx, "SELECT g.id, g.groupname, c.lessons FROM groups as g JOIN class_schedule AS c ON g.id=c.group_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var sch models.GroupSchedule
		if err := rows.Scan(&sch.Id, &sch.GroupName, &sch.Lessons); err != nil {
			return nil, err
		}
		schedules = append(schedules, sch)
	}
	return schedules, nil
}
