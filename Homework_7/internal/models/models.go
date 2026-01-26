package models

import "time"

type Student struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	GroupId   int       `json:"group_id"`
	GroupName string    `json:"group_name"`
	Gender    string    `json:"gender"`
	Birthd    time.Time `json:"birth_date"`
}

type StudentReq struct {
	Name    string    `json:"name"`
	GroupId int       `json:"group_id"`
	Gender  string    `json:"gender"`
	Birthd  time.Time `json:"birth_date"`
}

type Instructor struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Birthd    time.Time `json:"birth_date"`
	TitleName int       `json:"title_name"`
}

type Title struct {
	Id    int    `json:"id"`
	Title string `json:"title"` // instructor, assistant professor, associate professor, full professor (or just professor)
}

type GroupSchedule struct {
	Id         int    `json:"id"`
	GroupId    int    `json:"group_id"`
	Subject    string `json:"subject"`
	Start_Time string `json:"start_time"`
	End_Time   string `json:"end_time"`
}

type GroupReq struct {
	Groupname string `json:"groupname"`
}

type Attendance struct {
	Id        int       `json:"id"`
	SubjectId int       `json:"subject_id"`
	VisitDay  time.Time `json:"visit_day"`
	Visited   bool      `json:"visited"`
	StudentId int       `json:"student_id"`
}

type Grade struct {
	Id        int     `json:"id"`
	NumGrade  float64 `json:"grade"`
	Credits   int     `json:"credits"`
	StudentId int     `json:"student_id"`
	SubjectId int     `json:"subject_id"`
}

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"-"`
}

type AuthenticationReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
