package controllers

import (
	"mini_project/package/models"
	"mini_project/package/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all users
func GetAllPenjadwalanController(c echo.Context) error {
	ListPenjadwalan, err := db.GetAllPenjadwalan()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all penjadwalan",
		"Penjadwalan":   ListPenjadwalan,
	})
}

// get user by id
func GetPenjadwalanController(c echo.Context) error {
	Penjadwalan, err := db.GetSinglePenjadwalanById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get penjadwalan",
		"Penjadwalan":    Penjadwalan,
	})
}

// create new user
func CreatePenjadwalanController(c echo.Context) error {
	id := c.Param("id")
	Penjadwalan := models.Penjadwalan{}
	c.Bind(&Penjadwalan)
	err := db.CreateNewPenjadwalan(Penjadwalan)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create penjadwalan with id '" + id + "'",
	})
}

// delete user by id
func DeletePenjadwalanController(c echo.Context) error {
	id := c.Param("id")
	err := db.DeletePenjadwalanById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete penjadwalan with id '" + id + "'",
	})
}

// update user by id
func UpdatePenjadwalanController(c echo.Context) error {
	Penjadwalan := models.Penjadwalan{}
	c.Bind(&Penjadwalan)
	err := db.UpdatePenjadwalanById(c.Param("id"), Penjadwalan)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update penjadwalan",
		"Penjadwalan":    Penjadwalan,
	})
}