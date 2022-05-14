package controllers

import (
	"mini_project/package/models"
	"mini_project/package/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all users
func GetAllKaryawanController(c echo.Context) error {
	ListKaryawan, err := db.GetAllKaryawan()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all karyawan",
		"Karyawan":   ListKaryawan,
	})
}

// get user by id
func GetKaryawanController(c echo.Context) error {
	Karyawan, err := db.GetSingleKaryawanById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get karyawan",
		"Karyawan":    Karyawan,
	})
}

// create new user
func CreateKaryawanController(c echo.Context) error {
	id := c.Param("id")
	Karyawan := models.Karyawan{}
	c.Bind(&Karyawan)
	err := db.CreateNewKaryawan(Karyawan)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create karyawan with id '" + id + "'",
	})
}

// delete user by id
func DeleteKaryawanController(c echo.Context) error {
	id := c.Param("id")
	err := db.DeleteKaryawanById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete karyawan with id '" + id + "'",
	})
}

// update user by id
func UpdateKaryawanController(c echo.Context) error {
	Karyawan := models.Karyawan{}
	c.Bind(&Karyawan)
	err := db.UpdateKaryawanById(c.Param("id"), Karyawan)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update karyawan",
		"Karyawan":    Karyawan,
	})
}