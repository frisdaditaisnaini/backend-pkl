package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) CreatePivotDivisionField(pivot_division_field model.Pivot_division_field) error {
	res := r.DB.Create(&pivot_division_field)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert new division field")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllDivisionField(division_id int) []model.List_division_field {
	res := []model.List_division_field{}
	r.DB.Model(&model.Pivot_division_field{}).Select("study_fields.id,study_fields.name").
		Joins("JOIN divisions on divisions.id = pivot_division_fields.division_id").
		Joins("JOIN study_fields on study_fields.id = pivot_division_fields.study_field_id").
		Where("pivot_division_fields.division_id = ?", division_id).
		Scan(&res)

	return res
}

func (r *repositoryMysqlLayer) DeleteOnePivotDivisionField(division_id, study_field_id int) error {
	res := r.DB.Unscoped().Where("division_id = ? AND study_field_id = ?", division_id, study_field_id).Delete(&model.Pivot_division_field{})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete, division and study field not found")
	}

	return nil
}

func (r *repositoryMysqlLayer) CheckPivotDivisionByID(id int) (pivot_division_field model.Pivot_division_field, err error) {
	res := r.DB.Where("division_id = ?", id).Find(&pivot_division_field)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("division not found")
	}

	return
}

func (r *repositoryMysqlLayer) CheckPivotStudyFieldByID(id int) (pivot_division_field model.Pivot_division_field, err error) {
	res := r.DB.Where("study_field_id = ?", id).Find(&pivot_division_field)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("study field not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetAllDivisionStudyField() []model.Pivot_division_field {
	res := []model.Pivot_division_field{}
	r.DB.Find(&res)

	return res
}

func (r *repositoryMysqlLayer) GetDivisionOnPivot(id int) []model.Pivot_division_field {
	res := []model.Pivot_division_field{}
	r.DB.Where("division_id = ?", id).Find(&res)

	return res
}

func (r *repositoryMysqlLayer) CheckDivisonField(id_division, id_study_field int) (pivot_division_field model.Pivot_division_field, err error) {
	res := r.DB.Where("division_id = ? AND study_field_id = ?", id_division, id_study_field).Find(&pivot_division_field)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("division and study field not found")
	}

	return
}

func (r *repositoryMysqlLayer) DeleteAllDivisionField(division_id int) error {
	res := r.DB.Unscoped().Where("division_id = ?", division_id).Delete(&model.Pivot_division_field{})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete, division on pivot not found")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteAllStudyField(study_field_id int) error {
	res := r.DB.Unscoped().Where("study_field_id = ?", study_field_id).Delete(&model.Pivot_division_field{})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete, study field on pivot not found")
	}

	return nil
}
