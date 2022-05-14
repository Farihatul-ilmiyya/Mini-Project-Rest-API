package controllers

import (
	"mini_project/package/models"
	"mini_project/package/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all users
func GetAllWilayahKantorController(c echo.Context) error {
	ListWilayahKantor, err := db.GetAllWilayahKantor()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all wilayah kantor",
		"WilayahKantor":   ListWilayahKantor,
	})
}

// get user by id
func GetWilayahKantorController(c echo.Context) error {
	WilayahKantor, err := db.GetSingleWilayahKantorById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get wilayah kantor",
		"WilayahKantor":    WilayahKantor,
	})
}

// create new user
func CreateWilayahKantorController(c echo.Context) error {
	id := c.Param("id")
	WilayahKantor := models.WilayahKantor{}
	c.Bind(&WilayahKantor)
	err := db.CreateNewWilayahKantor(WilayahKantor)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create wilayah kantor with id '" + id + "'",
	})
}

// delete user by id
func DeleteWilayahKantorController(c echo.Context) error {
	id := c.Param("id")
	err := db.DeleteWilayahKantorById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete wilayah kantor with id '" + id + "'",
	})
}

// update user by id
func UpdateWilayahKantorController(c echo.Context) error {
	WilayahKantor := models.WilayahKantor{}
	c.Bind(&WilayahKantor)
	err := db.UpdateWilayahKantorById(c.Param("id"), WilayahKantor)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update wilayah kantor",
		"WilayahKantor":    WilayahKantor,
	})
}