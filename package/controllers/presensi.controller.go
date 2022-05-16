package controllers

import (
	"mini-project/package/models"
	"mini-project/package/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all presensi
func GetAllPresensiController(c echo.Context) error {
	ListPresensi, err := repository.GetAllPresensi()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all presensi",
		"Presensi":   ListPresensi,
	})
}

// get presensi by id
func GetPresensiController(c echo.Context) error {
	Presensi, err := repository.GetPresensiById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get presensi",
		"Presensi":    Presensi,
	})
}

// create new presensi
func CreatePresensiController(c echo.Context) error {
	id := c.Param("id")
	Presensi := models.Presensi{}
	c.Bind(&Presensi)
	err := repository.CreateNewPresensi(Presensi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create presensi with id '" + id + "'",
	})
}

// delete presensi by id
func DeletePresensiController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeletePresensiById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete presensi with id '" + id + "'",
	})
}

// update presensi by id
func UpdatePresensiController(c echo.Context) error {
	Presensi := models.Presensi{}
	c.Bind(&Presensi)
	err := repository.UpdatePresensiById(c.Param("id"), Presensi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update presensi",
		"Presensi":    Presensi,
	})
}