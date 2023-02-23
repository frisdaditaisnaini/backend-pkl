package model

import (
	"time"
)

type Trainee struct {
	ID                 uint       `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Name               string     `gorm:"not null" json:"name" form:"name"`
	Trainee_Student_id string     `gorm:"not null" json:"trainee_student_id" form:"trainee_student_id"`
	Email              string     `gorm:"not null" json:"email" form:"email"`
	Jurusan            string     `gorm:"not null" json:"jurusan" form:"jurusan"`
	Gender             string     `gorm:"not null" json:"gender" form:"gender"`
	Phone              string     `gorm:"not null" json:"phone" form:"phone"`
	CreatedAt          time.Time  `json:"created_at" form:"created_at"`
	SubmissionID       uint       `json:"submission_id" form:"submission_id"`
	Submission         Submission `json:"submission" form:"submission" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
