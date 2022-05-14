package controllers

import (
	"mini_project/package/models"
	"mini_project/package/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all users
func GetAllDivisiController(c echo.Context) error {
	ListDivisi, err := db.GetAllDivisi()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all divisi",
		"Divisi":   ListDivisi,
	})
}

// get user by id
func GetDivisiController(c echo.Context) error {
	Divisi, err := db.GetSingleDivisiById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get divisi",
		"Divisi":    Divisi,
	})
}

// create new user
func CreateDivisiController(c echo.Context) error {
	id := c.Param("id")
	Divisi := models.Divisi{}
	c.Bind(&Divisi)
	err := db.CreateNewDivisi(Divisi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create divisi with id '" + id + "'",
	})
}

// delete user by id
func DeleteDivisiController(c echo.Context) error {
	id := c.Param("id")
	err := db.DeleteDivisiById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete divisi with id '" + id + "'",
	})
}

// update user by id
func UpdateDivisiController(c echo.Context) error {
	Divisi := models.Divisi{}
	c.Bind(&Divisi)
	err := db.UpdateDivisiById(c.Param("id"), Divisi)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update divisi",
		"Divisi":    Divisi,
	})
}