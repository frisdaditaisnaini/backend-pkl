package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateSubmissionController(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error upload file",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error open file",
		})
	}

	if file.Size >= 2200000 {
		return c.JSON(400, map[string]interface{}{
			"messages": "max file is 1 MB",
		})
	}

	filebyte, _ := io.ReadAll(src)
	filetype := http.DetectContentType(filebyte)
	if filetype != "application/pdf" {
		return c.JSON(400, map[string]interface{}{
			"messages": "file is not .pdf",
		})
	}

	defer src.Close()

	submission := model.Submission{}
	if err := c.Bind(&submission); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	filename := "uploads/" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
	submission.SubmissionPathFile = filename

	err = os.WriteFile(filename, filebyte, 0777)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error write file",
		})
		//server_filename := strings.ReplaceAll(filename, "..", "")
		//server_filename_path := fmt.Sprintf("%s%s", "~/pkl", server_filename)
		//err = os.WriteFile(os.ExpandEnv(filename), filebyte, 0777)
		//submission.SubmissionPathFile = os.ExpandEnv(filename)
		//if err != nil {
		//	return c.JSON(400, map[string]interface{}{
		//		"messages": err.Error(),
		//	})
		//}
	}

	res, err := ce.Svc.CreateSubmissionService(submission)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":        "success",
		"code_submission": res,
	})
}

func (ce *EchoController) GetAllSubmissionController(c echo.Context) error {

	submissions := ce.Svc.GetAllSubmissionService()

	return c.JSON(200, map[string]interface{}{
		"messages":   "success",
		"submission": submissions,
	})
}

func (ce *EchoController) GetOneSubmissionController(c echo.Context) error {
	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	}

	return c.JSON(200, map[string]interface{}{
		"messages":   "success",
		"submission": res,
	})
}

func (ce *EchoController) GetFileSubmissionController(c echo.Context) error {
	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "file submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "file submission not found",
			})
		}
	}

	name_file := fmt.Sprintf("%s%s", res.CodeSubmission, ".pdf")

	return c.Attachment(res.SubmissionPathFile, name_file)
}

func (ce *EchoController) AcceptSubmissionController(c echo.Context) error {
	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error upload file",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error open file",
		})
	}

	if file.Size >= 2200000 {
		return c.JSON(400, map[string]interface{}{
			"messages": "max file is 1 MB",
		})
	}

	filebyte, _ := io.ReadAll(src)
	filetype := http.DetectContentType(filebyte)
	if filetype != "application/pdf" {
		return c.JSON(400, map[string]interface{}{
			"messages": "file is not .pdf",
		})
	}

	defer src.Close()

	filename := "uploads/" + res.CodeSubmission + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "_diterima_respon.pdf"
	res.ResponPathFile = filename
	res.Status = "Diterima"

	err = os.WriteFile(filename, filebyte, 0777)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error write new file respon",
		})
	}

	err = ce.Svc.UpdateSubmissionByIDService(int(res.ID), res)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": "no change",
		})
	}

	message := fmt.Sprintf("%s%s%s", "Pengajuan dengan kode ", res.CodeSubmission, " telah diterima !")

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"detail":   message,
	})
}

func (ce *EchoController) RejectSubmissionController(c echo.Context) error {
	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error upload file",
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error open file",
		})
	}

	if file.Size >= 2200000 {
		return c.JSON(400, map[string]interface{}{
			"messages": "max file is 1 MB",
		})
	}

	filebyte, _ := io.ReadAll(src)
	filetype := http.DetectContentType(filebyte)
	if filetype != "application/pdf" {
		return c.JSON(400, map[string]interface{}{
			"messages": "file is not .pdf",
		})
	}

	defer src.Close()

	filename := "uploads/" + res.CodeSubmission + "_" + strconv.FormatInt(time.Now().Unix(), 10) + "_ditolak_respon.pdf"
	res.ResponPathFile = filename
	res.Status = "Ditolak"

	err = os.WriteFile(filename, filebyte, 0777)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error write new file respon",
		})
	}

	err = ce.Svc.UpdateSubmissionByIDService(int(res.ID), res)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": "no change",
		})
	}

	message := fmt.Sprintf("%s%s%s", "Pengajuan dengan kode ", res.CodeSubmission, " telah ditolak !")

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"detail":   message,
	})
}

func (ce *EchoController) CancelSubmissionController(c echo.Context) error {
	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "submission not found",
			})
		}
	}
	res.Status = "Dibatalkan"
	res.ResponPathFile = " "

	err = ce.Svc.UpdateSubmissionByIDService(int(res.ID), res)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": "no change",
		})
	}

	message := fmt.Sprintf("%s%s%s", "Pengajuan dengan kode ", res.CodeSubmission, " telah dibatalkan !")

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"detail":   message,
	})
}

func (ce *EchoController) GetFileResponSubmissionController(c echo.Context) error {
	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "file submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "file submission not found",
			})
		}
	}

	name_file := fmt.Sprintf("%s%s", res.CodeSubmission, "_respon.pdf")

	return c.Attachment(res.ResponPathFile, name_file)
}

func (ce *EchoController) GetAllSubmissionByStatusController(c echo.Context) error {
	param := c.Param("status")
	submissions := ce.Svc.GetAllSubmissionByStatusService(param)

	return c.JSON(200, map[string]interface{}{
		"messages":   "success",
		"submission": submissions,
	})
}

func (ce *EchoController) DeleteSubmissionController(c echo.Context) error {

	var res model.Submission
	var err error
	param := c.Param("id_code")
	if strings.Contains(param, "P-") {
		res, err = ce.Svc.GetSubmissionByCodeSubmissionService(param)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "file submission not found",
			})
		}
	} else {
		id_int, _ := strconv.Atoi(param)
		res, err = ce.Svc.GetSubmissionByIDService(id_int)
		if err != nil {
			return c.JSON(404, map[string]interface{}{
				"messages": "file submission not found",
			})
		}
	}

	err = ce.Svc.DeleteSubmissionByIDService(int(res.ID))
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(204, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoController) ExportSubmissionToExcelController(c echo.Context) error {
	submissions := ce.Svc.GetAllSubmissionService()

	var err error

	file := excelize.NewFile()

	file.SetCellValue("Sheet1", "A1", "ID")
	file.SetCellValue("Sheet1", "B1", "Kode Pengajuan")
	file.SetCellValue("Sheet1", "C1", "Nama")
	file.SetCellValue("Sheet1", "D1", "Email")
	file.SetCellValue("Sheet1", "E1", "Asal Sekolah")
	file.SetCellValue("Sheet1", "F1", "Jumlah Anggota")
	file.SetCellValue("Sheet1", "G1", "Status")
	file.SetCellValue("Sheet1", "H1", "Submission Path File")
	file.SetCellValue("Sheet1", "I1", "Respon Path File")
	file.SetCellValue("Sheet1", "J1", "Tanggal Mulai")
	file.SetCellValue("Sheet1", "K1", "Tanggal Selesai")
	file.SetCellValue("Sheet1", "L1", "ID Divisi")
	file.SetCellValue("Sheet1", "M1", "Nama Divisi")
	file.SetCellValue("Sheet1", "N1", "ID Bidang Keilmuan")
	file.SetCellValue("Sheet1", "O1", "Nama Bidang Keilmuan")
	file.SetCellValue("Sheet1", "P1", "Created At")

	err = file.AutoFilter("Sheet1", "A1", "P1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	for i := 0; i < len(submissions); i++ {
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), submissions[i].ID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), submissions[i].CodeSubmission)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), submissions[i].Name)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), submissions[i].Email)
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), submissions[i].SchoolOrigin)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), submissions[i].TotalTrainee)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), submissions[i].Status)
		file.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), submissions[i].SubmissionPathFile)
		file.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2), submissions[i].ResponPathFile)
		file.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+2), submissions[i].StartDate)
		file.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+2), submissions[i].EndDate)
		file.SetCellValue("Sheet1", fmt.Sprintf("L%d", i+2), submissions[i].DivisionID)
		file.SetCellValue("Sheet1", fmt.Sprintf("M%d", i+2), submissions[i].Division.Name)
		file.SetCellValue("Sheet1", fmt.Sprintf("N%d", i+2), submissions[i].Study_fieldID)
		file.SetCellValue("Sheet1", fmt.Sprintf("O%d", i+2), submissions[i].Study_field.Name)
		file.SetCellValue("Sheet1", fmt.Sprintf("P%d", i+2), submissions[i].CreatedAt)
	}

	err = file.SaveAs("uploads/export.xlsx")
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": "error save file",
		})
	}

	return c.Attachment("uploads/export.xlsx", "export.xlsx")
}
