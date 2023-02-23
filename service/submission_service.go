package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	nanoid "github.com/aidarkhanov/nanoid/v2"
)

func (s *svc) CreateSubmissionService(submission model.Submission) (string, error) {
	if submission.Name == "" || submission.Email == "" || submission.SchoolOrigin == "" {
		return "", fmt.Errorf("insert submission error")
	}

	_, err := s.repo.CheckDivisonField(int(submission.DivisionID), int(submission.Study_fieldID))
	if err != nil {
		return "", fmt.Errorf("can't insert submission, division and study field not match")
	}

	division, _ := s.repo.GetDivisionByID(int(submission.DivisionID))
	if int(*division.Quota) < submission.TotalTrainee {
		return "", fmt.Errorf("can't insert submission, exceed max quota")
	}

	code, _ := nanoid.New()
	submission.CodeSubmission = code

	err = s.repo.CreateSubmission(submission)
	if err != nil {
		return "", err
	}

	res, err := s.repo.GetSubmissionByCodeSubmission(submission.CodeSubmission)
	if err != nil {
		return "", err
	}

	dt := res.CreatedAt.Format("01-02")
	format_dt := strings.ReplaceAll(dt, "-", "")
	int_dt, _ := strconv.Atoi(format_dt)
	res.CodeSubmission = fmt.Sprintf("%s%02d%s%02d%02d%s%d", "P-", int_dt, "-", int(res.DivisionID), int(res.Study_fieldID), "-", 1000+int(res.ID))
	err = s.repo.UpdateSubmissionByID(int(res.ID), res)
	if err != nil {
		err = s.repo.UpdateSubmissionByCodeSubmission(code, res)
		if err != nil {
			return "", err
		}
		return "", err
	}

	return res.CodeSubmission, nil
}

func (s *svc) GetAllSubmissionService() []model.Submission {
	return s.repo.GetAllSubmission()
}

func (s *svc) GetSubmissionByCodeSubmissionService(code_submission string) (model.Submission, error) {
	return s.repo.GetSubmissionByCodeSubmission(code_submission)
}

func (s *svc) GetSubmissionByIDService(id int) (model.Submission, error) {
	return s.repo.GetSubmissionByID(id)
}

func (s *svc) UpdateSubmissionByIDService(id int, submission model.Submission) error {
	return s.repo.UpdateSubmissionByID(id, submission)
}

func (s *svc) GetAllSubmissionByStatusService(status string) []model.Submission {
	return s.repo.GetAllSubmissionByStatus(status)
}

func (s *svc) DeleteSubmissionByIDService(id int) error {
	return s.repo.DeleteSubmissionByID(id)
}
