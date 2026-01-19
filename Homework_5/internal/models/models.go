package models

import "time"

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

type Attendance struct {
	SubjectName      string    `json:"subject_name"`
	SubjectId        int       `json:"subject_id"`
	VisitDay         time.Time `json:"visit_day"`
	Visited          bool      `json:"visited"`
	StudentFirstname string    `json:"student_firstname"`
	StudentSurname   string    `json:"student_surname"`
	StudentId        int       `json:"student_id"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type AuthenticationReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
