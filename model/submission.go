package model

import "time"

type Submission struct {
	ID                 uint        `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CodeSubmission     string      `gorm:"unique;not null" json:"code_submission" form:"code_submission"`
	Name               string      `gorm:"not null" json:"name" form:"name"`
	Email              string      `gorm:"unique;not null" json:"email" form:"email"`
	SchoolOrigin       string      `gorm:"not null" json:"school_origin" form:"school_origin"`
	TotalTrainee       int         `gorm:"not null" json:"total_trainee" form:"total_trainee"`
	Status             string      `gorm:"default:Diproses" json:"status" form:"status"`
	SubmissionPathFile string      `json:"path_file" form:"path_file"`
	ResponPathFile     string      `json:"respon_path_file" form:"respon_path_file"`
	StartDate          time.Time   `gorm:"not null" json:"start_date" form:"start_date"`
	EndDate            time.Time   `gorm:"not null" json:"end_date" form:"end_date"`
	CreatedAt          time.Time   `json:"created_at" form:"created_at"`
	DivisionID         uint        `json:"division_id" form:"division_id"`
	Division           Division    `json:"division" form:"division" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Study_fieldID      uint        `json:"study_field_id" form:"study_field_id"`
	Study_field        Study_field `json:"study_field" form:"study_field" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
