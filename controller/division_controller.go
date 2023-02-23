package controller

import (
	"strconv"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateDivisionController(c echo.Context) error {

	/* username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	} */

	division := model.Division{}
	if err := c.Bind(&division); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err := ce.Svc.CreateDivisionService(division)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(201, map[string]interface{}{
		"messages":      "success",
		"division_name": division.Name,
	})
}

func (ce *EchoController) GetAllDivisionController(c echo.Context) error {

	divisions := ce.Svc.GetAllDivisionService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"division": divisions,
	})
}

func (ce *EchoController) GetOneDivisionController(c echo.Context) error {
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	res, err := ce.Svc.GetDivisionByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "Division not found",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"division": res,
	})
}

func (ce *EchoController) UpdateDivisionController(c echo.Context) error {

	/* username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	} */

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	division := model.Division{}
	if err := c.Bind(&division); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err := ce.Svc.UpdateDivisionByIDService(id_int, division)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": "no id found or no change",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "updated",
	})
}

func (ce *EchoController) DeleteDivisionController(c echo.Context) error {
	/* username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden",
		})
	} */

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	err := ce.Svc.DeleteDivisionByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(204, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoController) GetChartAllDivisionController(c echo.Context) error {

	divisions := ce.Svc.GetChartAllDivisionService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"division": divisions,
	})
}
