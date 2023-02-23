package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) GetAdminByUsername(username string) (admin model.Admin, err error) {
	res := r.DB.Where("username = ?", username).Find(&admin)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("admin not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateAdminByID(id int, admin model.Admin) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&admin)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update admin")
	}

	return nil
}
