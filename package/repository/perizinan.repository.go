package repository

import (
	"fmt"
	"mini-project/package/db"
	"mini-project/package/models"
)

func GetAllPerizinan() ([]models.Perizinan, error) {
	var ListPerizinan []models.Perizinan

	err := database.DB.Find(&ListPerizinan).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListPerizinan, err
}

func GetPerizinanById(id string) (models.Perizinan, error) {
	var Perizinan models.Perizinan
	err := database.DB.First(&Perizinan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return Perizinan, err
}

func DeletePerizinanById(id string) error {
	err := database.DB.Delete(&models.Perizinan{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewPerizinan(Perizinan models.Perizinan) error {
	err := database.DB.Save(&Perizinan).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdatePerizinanById(id string, Perizinan models.Perizinan) error {
	err := database.DB.Model(&Perizinan).Where("id = ?", id).Updates(Perizinan).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}