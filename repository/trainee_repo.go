package repository

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"gorm.io/gorm/clause"
)

func (r *repositoryMysqlLayer) CreateTrainee(trainee model.Trainee) error {
	res := r.DB.Create(&trainee)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert trainee repo")
	}

	return nil
}

func (r *repositoryMysqlLayer) CountTrainee(submission_id int) int {
	var count int64
	r.DB.Model(&model.Trainee{}).Where("submission_id = ?", submission_id).Count(&count)

	return int(count)
}

func (r *repositoryMysqlLayer) GetAllTraineeByIDSubmission(id int) []model.Trainee {
	trainees := []model.Trainee{}
	r.DB.Where("submission_id = ?", id).Preload("Submission.Division").Preload("Submission.Study_field").Preload(clause.Associations).Find(&trainees)

	return trainees
}

func (r *repositoryMysqlLayer) GetAllTrainee() []model.Trainee {
	trainees := []model.Trainee{}
	r.DB.Preload("Submission.Division").Preload("Submission.Study_field").Preload(clause.Associations).Find(&trainees)

	return trainees
}

func (r *repositoryMysqlLayer) GetOneTraineeByID(id int) (trainee model.Trainee, err error) {
	res := r.DB.Where("id = ?", id).Preload("Submission.Division").Preload("Submission.Study_field").Preload(clause.Associations).Find(&trainee)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("id trainee not found")
	}

	return
}
