package controller

import (
	"strconv"
	"strings"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateTraineeController(c echo.Context) error {

	var submission model.Submission
	var err error
	param := c.Param("id_code_submission")
	if strings.Contains(param, "P-") {
		submission, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		submission, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	}

	trainee := model.Trainee{}
	if err := c.Bind(&trainee); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	trainee.SubmissionID = submission.ID
	count := ce.Svc.CountTraineeService(int(submission.ID))
	if count >= submission.TotalTrainee {
		return c.JSON(400, map[string]interface{}{
			"messages": "total trainee is full for this submission",
		})
	}

	err = ce.Svc.CreateTraineeService(trainee)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	trainee.Submission = submission

	return c.JSON(200, map[string]interface{}{
		"messages":   "success",
		"submission": submission.CodeSubmission,
		"trainee":    trainee,
	})
}

func (ce *EchoController) GetAllTraineeByIDSubmissionController(c echo.Context) error {
	var submission model.Submission
	var err error
	param := c.Param("id_code_submission")
	if strings.Contains(param, "P-") {
		submission, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		submission, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	}

	trainees := ce.Svc.GetAllTraineeByIDSubmissionService(int(submission.ID))

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"trainee":  trainees,
	})
}

func (ce *EchoController) GetAllTraineeController(c echo.Context) error {

	trainees := ce.Svc.GetAllTraineeService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"trainee":  trainees,
	})
}

func (ce *EchoController) GetOneTraineeController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	trainee, err := ce.Svc.GetOneTraineeByIDService(id)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "trainee not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"trainee":  trainee,
	})
}
