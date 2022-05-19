package controllers

import (
	"mini-project/package/helpers"
	"mini-project/package/models"
	"net/http"
	"mini-project/package/repository"
	"github.com/labstack/echo/v4"
)

// func CheckLogin(c echo.Context) error {
// 	username := c.FormValue("username")
// 	password := c.FormValue("password")

// 	res, err := models.CheckLogin(username, password)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"message": err.Error(),
// 		})
// 	}

// 	if !res {
// 		return echo.ErrUnauthorized
		
// 	}
// 	return c.String(http.StatusOK, "berhasil login")
// }

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)
	return c.JSON(http.StatusOK, hash)
}

func CreateUserController(c echo.Context) error {
	NewUser := models.User{}
	c.Bind(&NewUser)
	err := repository.CreateNewUser(NewUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Membuat Data User",
	})

}

// func CreateLoginController(c echo.Context) error {
// 	NewUser := models.User{}
// 	c.Bind(&NewUser)
// 	err := repository.CreateNewUser(NewUser)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "Berhasil Membuat Data User",
// 	})

// }


//user controller
// package controllers

// import (
// 	"mini-project/package/models"
// 	"mini-project/package/repository"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// 	// get all users
// func GetAllUserController(c echo.Context) error {
// 	ListUser, err := repository.GetAllUser()
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get all user",
// 		"User":   ListUser,
// 	})
// }

// // get user by id
// func GetUserController(c echo.Context) error {
// 	User, err := repository.GetUserById(c.Param("id"))
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success get user",
// 		"User":    User,
// 	})
// }

// // create new user
// func CreateUserController(c echo.Context) error {
// 	id := c.Param("id")
// 	User := models.User{}
// 	c.Bind(&User)
// 	err := repository.CreateNewUser(User)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success create user with id '" + id + "'",
// 	})
// }

// // delete user by id
// func DeleteUserController(c echo.Context) error {
// 	id := c.Param("id")
// 	err := repository.DeleteUserById(id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success delete user with id '" + id + "'",
// 	})
// }

// // update user by id
// func UpdateUserController(c echo.Context) error {
// 	User := models.User{}
// 	c.Bind(&User)
// 	err := repository.UpdateUserById(c.Param("id"), User)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"message": "success update user",
// 		"User":    User,
// 	})
// }