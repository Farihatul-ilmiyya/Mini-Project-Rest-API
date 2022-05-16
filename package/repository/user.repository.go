package repository

import (
	"fmt"
	"mini-project/package/db"
	"mini-project/package/models"
)

func GetAllUser() ([]models.User, error) {
	var ListUser []models.User

	err := database.DB.Find(&ListUser).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListUser, err
}

func GetUserById(id string) (models.User, error) {
	var User models.User
	err := database.DB.First(&User, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return User, err
}

func DeleteUserById(id string) error {
	err := database.DB.Delete(&models.User{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewUser(User models.User) error {
	err := database.DB.Save(&User).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateUserById(id string, User models.User) error {
	err := database.DB.Model(&User).Where("id = ?", id).Updates(User).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}