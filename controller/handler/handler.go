package handler

import (
	"fmt"
	"os"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/config"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/controller"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/database"

	m "github.com/PKL-Angkasa-Pura-I/backend-pkl/middleware"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/repository"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewService(repo, conf)

	cont := controller.EchoController{
		Svc: svc,
	}

	e.GET("/pkl_v1/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	e.GET("/pkl_v1/workdir", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": mydir,
		})
	})

	e.POST("/pkl_v1/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	api := e.Group("/pkl_v1", middleware.CORS())

	m.LogMiddleware(e)
	api.POST("/admins/login", cont.LoginAdminController)

	api.POST("/admins/passwords", cont.ChangePassAdminController)

	api.POST("/divisions", cont.CreateDivisionController)
	api.GET("/divisions", cont.GetAllDivisionController)

	api.GET("/divisions/:id", cont.GetOneDivisionController)
	api.PUT("/divisions/:id", cont.UpdateDivisionController)
	api.DELETE("/divisions/:id", cont.DeleteDivisionController)

	api.POST("/study_fields", cont.CreateStudyFieldController)
	api.GET("/study_fields", cont.GetAllStudyFieldController)

	api.GET("/study_fields/:id", cont.GetOneStudyFieldController)
	api.PUT("/study_fields/:id", cont.UpdateStudyFieldController)
	api.DELETE("/study_fields/:id", cont.DeleteStudyFieldController)

	api.GET("/list_division_fields", cont.GetAllDivisionStudyFieldController)

	api.POST("/list_division_fields/:id_division", cont.CreatePivotDivisionFieldController)
	api.GET("/list_division_fields/:id_division", cont.GetAllDivisionFieldController)
	api.DELETE("/list_division_fields/:id_division", cont.DeleteOnePivotDivisionFieldController)

	api.POST("/submissions", cont.CreateSubmissionController)
	api.GET("/submissions", cont.GetAllSubmissionController)

	api.GET("/submissions/export", cont.ExportSubmissionToExcelController)

	api.GET("/submissions/:id_code", cont.GetOneSubmissionController)
	api.DELETE("/submissions/:id_code", cont.DeleteSubmissionController)
	api.PUT("/submissions/:id_code/accept", cont.AcceptSubmissionController)
	api.PUT("/submissions/:id_code/reject", cont.RejectSubmissionController)
	api.PUT("/submissions/:id_code/cancel", cont.CancelSubmissionController)

	api.GET("/submissions/:id_code/download", cont.GetFileSubmissionController)
	api.GET("/submissions/:id_code/download/respon", cont.GetFileResponSubmissionController)

	api.GET("/submissions/filters/:status", cont.GetAllSubmissionByStatusController)

	api.GET("/trainees", cont.GetAllTraineeController)

	api.POST("/trainees/:id_code_submission", cont.CreateTraineeController)
	api.GET("/trainees/:id_code_submission", cont.GetAllTraineeByIDSubmissionController)

	api.GET("/trainees/details/:id", cont.GetOneTraineeController)

	api.GET("/charts/all_division", cont.GetChartAllDivisionController)

}
