package model

import "time"

type Admin struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id" form:"id"`
	Username  string    `gorm:"uniqueIndex;size:30" json:"username" form:"username"`
	Password  string    `json:"password" form:"password"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type AdminLogin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type AdminChangePass struct {
	OldPass string `json:"old_password" form:"old_password"`
	NewPass string `json:"new_password" form:"new_password"`
}
