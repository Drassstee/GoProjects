package storage

import (
	"context"
	"homework_7/internal/models"

	"github.com/jackc/pgx/v5"
)

type StudentStorage struct {
	DB *pgx.Conn
}

func New(db *pgx.Conn) *StudentStorage {
	return &StudentStorage{DB: db}
}

func (s *StudentStorage) GetInstructor(ctx context.Context, id string) (models.Instructor, error) {
	var it models.Instructor
	err := s.DB.QueryRow(ctx, "SELECT i.id, i.name, i.birth_date, i.gender, t.title FROM instructors AS i JOIN titles as t ON i.title_id=t.id WHERE i.id=$1", id).Scan(&it.Id, &it.Name, &it.Birthd, &it.Gender, &it.TitleName)
	return it, err
}

func (s *StudentStorage) GetStudent(ctx context.Context, id string) (models.Student, error) {
	var st models.Student
	err := s.DB.QueryRow(ctx, "SELECT s.id, s.name, s.birth_date, s.gender, g.id, g.groupname  FROM students AS s JOIN groups as g ON s.group_id=g.id WHERE s.id=$1", id).Scan(&st.Id, &st.Name, &st.Birthd, &st.Gender, &st.GroupId, &st.GroupName)
	return st, err
}

func (s *StudentStorage) InsertStudent(ctx context.Context, student models.StudentReq) (int, error) {
	var id int
	err := s.DB.QueryRow(ctx, "INSERT INTO students (name, birth_date, gender, group_id) VALUES ($1, $2, $3, $4) RETURNING id", student.Name, student.Birthd, student.Gender, student.GroupId).Scan(&id)
	return id, err
}

func (s *StudentStorage) InsertGroup(ctx context.Context, groupname string) (int, error) {
	var id int
	err := s.DB.QueryRow(ctx, "INSERT INTO groups (groupname) VALUES ($1) RETURNING id", groupname).Scan(&id)
	return id, err
}

func (s *StudentStorage) GetStudentGrades(ctx context.Context, id string) ([]models.Grade, error) {
	var grades []models.Grade
	rows, err := s.DB.Query(ctx, "SELECT g.id, g.grade, g.credits, g.student_id, g.subject_id FROM grades AS g WHERE g.id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var gr models.Grade
		if err := rows.Scan(&gr.Id, &gr.NumGrade, &gr.Credits, &gr.StudentId, &gr.SubjectId); err != nil {
			return nil, err
		}
		grades = append(grades, gr)
	}
	return grades, nil
}

func (s *StudentStorage) GetGroupSchedule(ctx context.Context, id string) (models.GroupSchedule, error) {
	var sch models.GroupSchedule
	err := s.DB.QueryRow(ctx, "SELECT c.id, g.id, c.subject, c.start_time, c.end_time FROM class_schedule AS c JOIN groups AS g ON c.group_id=g.id WHERE g.id=$1", id).Scan(&sch.Id, &sch.GroupId, &sch.Subject, &sch.Start_Time, &sch.End_Time)
	return sch, err
}

func (s *StudentStorage) GetAllSchedules(ctx context.Context) ([]models.GroupSchedule, error) {
	var schedules []models.GroupSchedule
	rows, err := s.DB.Query(ctx, "SELECT c.id, g.id, c.subject, c.start_time, c.end_time FROM class_schedule AS c JOIN groups AS g ON c.group_id=g.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var sch models.GroupSchedule
		if err := rows.Scan(&sch.Id, &sch.GroupId, &sch.Subject, &sch.Start_Time, &sch.End_Time); err != nil {
			return nil, err
		}
		schedules = append(schedules, sch)
	}
	return schedules, nil
}

func (s *StudentStorage) GetAttendanceStudent(ctx context.Context, id string) ([]models.Attendance, error) {
	var attendances []models.Attendance
	rows, err := s.DB.Query(ctx, "SELECT a.id, a.student_id, a.subject_id, a.visit_day, a.visit FROM attendance WHERE a.student_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a models.Attendance
		if err := rows.Scan(&a.Id, &a.StudentId, &a.SubjectId, &a.VisitDay, &a.Visited); err != nil {
			return nil, err
		}
		attendances = append(attendances, a)
	}
	return attendances, nil
}

func (s *StudentStorage) GetAttendanceSubject(ctx context.Context, id string) ([]models.Attendance, error) {
	var attendances []models.Attendance
	rows, err := s.DB.Query(ctx, "SELECT a.id, a.student_id, a.subject_id, a.visit_day, a.visit FROM attendance WHERE a.subject_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a models.Attendance
		if err := rows.Scan(&a.Id, &a.StudentId, &a.SubjectId, &a.VisitDay, &a.Visited); err != nil {
			return nil, err
		}
		attendances = append(attendances, a)
	}
	return attendances, nil
}

func (s *StudentStorage) InsertAttendance(ctx context.Context, attendace models.Attendance) (int, error) {
	var id int
	err := s.DB.QueryRow(ctx, "INSERT INTO attendance (student_id, subject_id, visit, visit_day) VALUES ($1, $2, $3, $4) RETURNING id", attendace.StudentId, attendace.SubjectId, attendace.Visited, attendace.VisitDay).Scan(&id)
	return id, err
}

func (s *StudentStorage) RegisterUser(ctx context.Context, email, password, role string) error {
	_, err := s.DB.Exec(ctx, "INSERT INTO users (email, password, role) VALUES ($1, $2, $3)", email, password, role)
	return err
}

func (s *StudentStorage) LoginUser(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := s.DB.QueryRow(ctx, "SELECT id, email, password FROM users WHERE email=$1", email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

func (s *StudentStorage) GetUserById(ctx context.Context, id int) (models.User, error) {
	var user models.User
	err := s.DB.QueryRow(ctx, "SELECT id, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Email)
	return user, err
}
