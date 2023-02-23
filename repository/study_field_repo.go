package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (r *repositoryMysqlLayer) CreateStudyField(study_field model.Study_field) error {
	res := r.DB.Create(&study_field)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert study field")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllStudyField() []model.Study_field {
	study_fields := []model.Study_field{}
	r.DB.Find(&study_fields)

	return study_fields
}

func (r *repositoryMysqlLayer) GetStudyFieldByID(id int) (study_field model.Study_field, err error) {
	res := r.DB.Where("id = ?", id).Find(&study_field)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("study field not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateStudyFieldByID(id int, study_field model.Study_field) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&study_field)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update study field")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteStudyFieldByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Study_field{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete study field, study field not found")
	}

	return nil
}
