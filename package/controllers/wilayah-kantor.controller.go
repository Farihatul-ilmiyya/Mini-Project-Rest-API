package controllers

import (
	"mini-project/package/models"
	"mini-project/package/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all wilayah kantor
func GetAllWilayahKantorController(c echo.Context) error {
	ListWilayahKantor, err := repository.GetAllWilayahKantor()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all wilayah kantor",
		"WilayahKantor":   ListWilayahKantor,
	})
}

// get wilayah kantor by id
func GetWilayahKantorController(c echo.Context) error {
	WilayahKantor, err := repository.GetWilayahKantorById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get wilayah kantor",
		"WilayahKantor":    WilayahKantor,
	})
}

// create new wilayah kantor
func CreateWilayahKantorController(c echo.Context) error {
	id := c.Param("id")
	WilayahKantor := models.WilayahKantor{}
	c.Bind(&WilayahKantor)
	err := repository.CreateNewWilayahKantor(WilayahKantor)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create wilayah kantor with id '" + id + "'",
	})
}

// delete wilayah kantor by id
func DeleteWilayahKantorController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteWilayahKantorById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete wilayah kantor with id '" + id + "'",
	})
}

// update wilayah kantor by id
func UpdateWilayahKantorController(c echo.Context) error {
	WilayahKantor := models.WilayahKantor{}
	c.Bind(&WilayahKantor)
	err := repository.UpdateWilayahKantorById(c.Param("id"), WilayahKantor)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update wilayah kantor",
		"WilayahKantor":    WilayahKantor,
	})
}