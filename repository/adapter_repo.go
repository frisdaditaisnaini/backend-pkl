package repository

import (
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/adapter"

	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) adapter.AdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}
