package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"gorm.io/gorm/clause"
)

func (r *repositoryMysqlLayer) CreateSubmission(submission model.Submission) error {
	res := r.DB.Create(&submission)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert submission")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetSubmissionByCodeSubmission(code_submission string) (submission model.Submission, err error) {
	res := r.DB.Where("code_submission = ?", code_submission).Preload(clause.Associations).Find(&submission)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("code submission not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetSubmissionByID(id int) (submission model.Submission, err error) {
	res := r.DB.Where("id = ?", id).Preload(clause.Associations).Find(&submission)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("id submission not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateSubmissionByID(id int, submission model.Submission) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&submission)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update submission by id")
	}

	return nil
}

func (r *repositoryMysqlLayer) UpdateSubmissionByCodeSubmission(code string, submission model.Submission) error {
	res := r.DB.Where("code_submission = ?", code).UpdateColumns(&submission)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update submission by code submission")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAllSubmission() []model.Submission {
	submissions := []model.Submission{}
	r.DB.Preload(clause.Associations).Find(&submissions)

	return submissions
}

func (r *repositoryMysqlLayer) GetAllSubmissionByStatus(status string) []model.Submission {
	submissions := []model.Submission{}
	r.DB.Where("status = ?", status).Preload(clause.Associations).Find(&submissions)

	return submissions
}

func (r *repositoryMysqlLayer) DeleteSubmissionByID(id int) error {
	res := r.DB.Unscoped().Delete(&model.Submission{}, id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete submission, id submission not found")
	}

	return nil
}
