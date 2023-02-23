package model

import "time"

type Study_field struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Name      string    `gorm:"unique;not null" json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}
