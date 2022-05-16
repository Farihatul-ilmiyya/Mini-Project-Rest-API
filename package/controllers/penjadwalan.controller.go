package controllers

import (
	"mini-project/package/models"
	"mini-project/package/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all penjadwalan
func GetAllPenjadwalanController(c echo.Context) error {
	ListPenjadwalan, err := repository.GetAllPenjadwalan()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all penjadwalan",
		"Penjadwalan":   ListPenjadwalan,
	})
}

// get penjadwalan by id
func GetPenjadwalanController(c echo.Context) error {
	Penjadwalan, err := repository.GetPenjadwalanById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get penjadwalan",
		"Penjadwalan":    Penjadwalan,
	})
}

// create new penjadwalan
func CreatePenjadwalanController(c echo.Context) error {
	id := c.Param("id")
	Penjadwalan := models.Penjadwalan{}
	c.Bind(&Penjadwalan)
	err := repository.CreateNewPenjadwalan(Penjadwalan)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create penjadwalan with id '" + id + "'",
	})
}

// delete penjadwalan by id
func DeletePenjadwalanController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeletePenjadwalanById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete penjadwalan with id '" + id + "'",
	})
}

// update penjadwalan by id
func UpdatePenjadwalanController(c echo.Context) error {
	Penjadwalan := models.Penjadwalan{}
	c.Bind(&Penjadwalan)
	err := repository.UpdatePenjadwalanById(c.Param("id"), Penjadwalan)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update penjadwalan",
		"Penjadwalan":    Penjadwalan,
	})
}