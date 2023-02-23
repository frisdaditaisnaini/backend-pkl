package model

import "time"

type Pivot_division_field struct {
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	CreatedAt     time.Time   `json:"created_at" form:"created_at"`
	DivisionID    uint        `json:"division_id" form:"division_id"`
	Division      Division    `json:"division" form:"division" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Study_fieldID uint        `json:"study_field_id" form:"study_field_id"`
	Study_field   Study_field `json:"study_field" form:"study_field" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type List_division_field struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"field_name" form:"field_name"`
}

type List_pivot struct {
	ID               uint     `json:"id" form:"id"`
	Quota            *uint    `json:"quota" form:"quota"`
	DivisionName     string   `json:"division_name" form:"division_name"`
	ListStudyFieldID []uint   `json:"list_study_field_id" form:"list_id_study_field_id"`
	ListStudyField   []string `json:"list_study_field" form:"list_study_field"`
}
