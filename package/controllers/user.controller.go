package controllers

import (
	"mini-project/package/models"
	"mini-project/package/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

	// get all users
func GetAllUserController(c echo.Context) error {
	ListUser, err := repository.GetAllUser()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all user",
		"User":   ListUser,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	User, err := repository.GetUserById(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"User":    User,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	id := c.Param("id")
	User := models.User{}
	c.Bind(&User)
	err := repository.CreateNewUser(User)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user with id '" + id + "'",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	err := repository.DeleteUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user with id '" + id + "'",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	User := models.User{}
	c.Bind(&User)
	err := repository.UpdateUserById(c.Param("id"), User)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"User":    User,
	})
}