package controllers

import (
	"mini-project/package/models"
	"mini-project/package/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all perizinan
func GetAllPerizinanController(c echo.Context) error {
	ListPerizinan, err := repository.GetAllPerizinan()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all perizinan",
		"Perizinan":   ListPerizinan,
	})
}

// get perizinan by id
func GetPerizinanController(c echo.Context) error {
	Perizinan, err := repository.GetPerizinanById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get perizinan",
		"Perizinan":    Perizinan,
	})
}

// create new perizinan
func CreatePerizinanController(c echo.Context) error {
	id := c.Param("id")
	Perizinan := models.Perizinan{}
	c.Bind(&Perizinan)
	err := repository.CreateNewPerizinan(Perizinan)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create perizinan with id '" + id + "'",
	})
}

// delete perizinan by id
func DeletePerizinanController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeletePerizinanById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete perizinan with id '" + id + "'",
	})
}

// update perizinan by id
func UpdatePerizinanController(c echo.Context) error {
	Perizinan := models.Perizinan{}
	c.Bind(&Perizinan)
	err := repository.UpdatePerizinanById(c.Param("id"), Perizinan)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update perizinan",
		"Perizinan":    Perizinan,
	})
}