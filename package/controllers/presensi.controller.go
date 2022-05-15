package controllers

import (
	"mini_project/package/models"
	"mini_project/package/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all users
func GetAllPresensiController(c echo.Context) error {
	ListPresensi, err := db.GetAllPresensi()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all presensi",
		"Presensi":   ListPresensi,
	})
}

// get user by id
func GetPresensiController(c echo.Context) error {
	Presensi, err := db.GetSinglePresensiById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get presensi",
		"Presensi":    Presensi,
	})
}

// create new user
func CreatePresensiController(c echo.Context) error {
	id := c.Param("id")
	Presensi := models.Presensi{}
	c.Bind(&Presensi)
	err := db.CreateNewPresensi(Presensi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create presensi with id '" + id + "'",
	})
}

// delete user by id
func DeletePresensiController(c echo.Context) error {
	id := c.Param("id")
	err := db.DeletePresensiById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete presensi with id '" + id + "'",
	})
}

// update user by id
func UpdatePresensiController(c echo.Context) error {
	Presensi := models.Presensi{}
	c.Bind(&Presensi)
	err := db.UpdatePresensiById(c.Param("id"), Presensi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update presensi",
		"Presensi":    Presensi,
	})
}