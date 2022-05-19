package controllers

import (
	"mini-project/package/models"
	"mini-project/package/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllDivisiController(c echo.Context) error {
	ListDivisi, err := repository.GetAllDivisi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all divisi",
		"Divisi":   ListDivisi,
	})
}

func GetDivisiController(c echo.Context) error {
	Divisi, err := repository.GetDivisiById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, Divisi)

}

func DeleteDivisiController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteDivisiById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete divisi with id '" + id + "'",
	})

}

func CreateDivisiController(c echo.Context) error {
	id := c.Param("id")
	NamaDivisi := models.Divisi{}
	c.Bind(&NamaDivisi)
	err := repository.CreateNewDivisi(NamaDivisi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success insert divisi with id '" + id + "'",
	})

}

func UpdateDivisiController(c echo.Context) error {
	NamaDivisi := models.Divisi{}
	c.Bind(&NamaDivisi)
	err := repository.UpdateDivisiById(c.Param("id"), NamaDivisi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, NamaDivisi)
}