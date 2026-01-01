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
