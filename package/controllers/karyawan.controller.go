package controllers

import (
	"mini-project/package/models"
	"mini-project/package/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllKaryawanController(c echo.Context) error {
	ListKaryawan, err := repository.GetAllKaryawan()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all Karyawan",
		"Karyawan":   ListKaryawan,
	})
}

func GetKaryawanController(c echo.Context) error {
	Karyawan, err := repository.GetKaryawanById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, Karyawan)

}

func DeleteKaryawanController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteKaryawanById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete karyawan with id '" + id + "'",
	})

}

func CreateKaryawanController(c echo.Context) error {
	id := c.Param("id")
	Karyawan := models.Karyawan{}
	c.Bind(&Karyawan)
	err := repository.CreateNewKaryawan(Karyawan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert karyawan with id '" + id + "'",
	})

}

func UpdateKaryawanController(c echo.Context) error {
	Karyawan := models.Karyawan{}
	c.Bind(&Karyawan)
	err := repository.UpdateKaryawanById(c.Param("id"), Karyawan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, Karyawan)

}