package service

import (
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
)

func (s *svc) CreatePivotDivisionFieldService(pivot_division_field model.Pivot_division_field) error {
	return s.repo.CreatePivotDivisionField(pivot_division_field)
}

func (s *svc) GetAllDivisionFieldService(division_id int) []model.List_division_field {
	return s.repo.GetAllDivisionField(division_id)
}

func (s *svc) DeleteOnePivotDivisionFieldService(division_id, study_field_id int) error {
	return s.repo.DeleteOnePivotDivisionField(division_id, study_field_id)
}

func (s *svc) GetAllDivisionStudyFieldService() []model.List_pivot {
	divisions := s.repo.GetAllDivision()
	len_division := len(divisions)

	res := make([]model.List_pivot, len_division)
	for i, division := range divisions {
		res[i].ID = division.ID
		res[i].Quota = division.Quota
		res[i].DivisionName = division.Name
		list_id_study_field := s.repo.GetDivisionOnPivot(int(division.ID))
		for _, list := range list_id_study_field {
			study_field, _ := s.repo.GetStudyFieldByID(int(list.Study_fieldID))
			res[i].ListStudyFieldID = append(res[i].ListStudyFieldID, study_field.ID)
			res[i].ListStudyField = append(res[i].ListStudyField, study_field.Name)
		}
	}
	return res
}
