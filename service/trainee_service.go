package service

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (s *svc) CreateTraineeService(trainee model.Trainee) error {
	if trainee.Name == "" || trainee.Email == "" || trainee.Trainee_Student_id == "" || trainee.Phone == "" || trainee.Jurusan == "" || trainee.Gender == "" {
		return fmt.Errorf("error insert trainee service")
	}
	return s.repo.CreateTrainee(trainee)
}

func (s *svc) CountTraineeService(submission_id int) int {
	return s.repo.CountTrainee(submission_id)
}

func (s *svc) GetAllTraineeByIDSubmissionService(id int) []model.Trainee {
	return s.repo.GetAllTraineeByIDSubmission(id)
}

func (s *svc) GetAllTraineeService() []model.Trainee {
	return s.repo.GetAllTrainee()
}

func (s *svc) GetOneTraineeByIDService(id int) (model.Trainee, error) {
	return s.repo.GetOneTraineeByID(id)
}
