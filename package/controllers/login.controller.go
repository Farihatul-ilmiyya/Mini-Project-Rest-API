package controllers

import (
	"net/http"
	"mini-project/package/helpers"
	"github.com/labstack/echo/v4"
)

func GenerateHashPassword(c echo.Context) error{
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)
	return c.JSON(http.StatusOK, hash)
}