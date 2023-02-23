package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) CreateDivision(division model.Division) error {
	res := r.DB.Create(&division)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert division")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllDivision() []model.Division {
	divisions := []model.Division{}
	r.DB.Find(&divisions)

	return divisions
}

func (r *repositoryMysqlLayer) GetDivisionByID(id int) (division model.Division, err error) {
	res := r.DB.Where("id = ?", id).Find(&division)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("division not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateDivisionByID(id int, division model.Division) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&division)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update division")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteDivisionByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Division{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete division, division not found")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetTotalAcceptedDivision(division_id int) int {
	//var count int64
	//r.DB.Model(&model.Submission{}).Where("division_id = ? AND status = ?", submission_id, "Diterima").Count(&count)

	type NResult struct {
		N int64
	}
	var n NResult
	r.DB.Table("submissions").Select("sum(total_trainee) as n").Where("division_id = ? AND status = ?", division_id, "Diterima").Scan(&n)

	return int(n.N)
}
