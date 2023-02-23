package model

import "time"

type Division struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Name      string    `gorm:"unique;not null" json:"name" form:"name"`
	Quota     *uint     `gorm:"default:0;not null" json:"quota" form:"quota"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type Chart_division struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Quota *uint  `json:"quota" form:"quota"`
	Total uint   `json:"total" form:"total"`
}
