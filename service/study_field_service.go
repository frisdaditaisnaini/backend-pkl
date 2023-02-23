package service

import (
	"fmt"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (s *svc) CreateStudyFieldService(study_field model.Study_field) error {
	if study_field.Name == "" {
		return fmt.Errorf("error insert study field")
	}
	return s.repo.CreateStudyField(study_field)
}

func (s *svc) GetAllStudyFieldService() []model.Study_field {
	return s.repo.GetAllStudyField()
}

func (s *svc) GetStudyFieldByIDService(id int) (model.Study_field, error) {
	return s.repo.GetStudyFieldByID(id)
}

func (s *svc) UpdateStudyFieldByIDService(id int, study_field model.Study_field) error {
	if study_field.Name == "" {
		return fmt.Errorf("error insert study field")
	}
	return s.repo.UpdateStudyFieldByID(id, study_field)
}

func (s *svc) DeleteStudyFieldByIDService(id int) error {
	s.repo.DeleteAllStudyField(id)

	return s.repo.DeleteStudyFieldByID(id)
}
