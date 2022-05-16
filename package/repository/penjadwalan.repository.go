package repository

import (
	"fmt"
	"mini-project/package/db"
	"mini-project/package/models"
)

func GetAllPenjadwalan() ([]models.Penjadwalan, error) {
	var ListPenjadwalan []models.Penjadwalan

	err := database.DB.Find(&ListPenjadwalan).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListPenjadwalan, err
}

func GetPenjadwalanById(id string) (models.Penjadwalan, error) {
	var Penjadwalan models.Penjadwalan
	err := database.DB.First(&Penjadwalan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return Penjadwalan, err
}

func DeletePenjadwalanById(id string) error {
	err := database.DB.Delete(&models.Penjadwalan{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewPenjadwalan(Penjadwalan models.Penjadwalan) error {
	err := database.DB.Save(&Penjadwalan).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdatePenjadwalanById(id string, Penjadwalan models.Penjadwalan) error {
	err := database.DB.Model(&Penjadwalan).Where("id = ?", id).Updates(Penjadwalan).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}