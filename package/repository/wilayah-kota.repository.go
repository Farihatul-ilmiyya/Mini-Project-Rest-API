package repository

import (
	"fmt"
	"mini-project/package/db"
	"mini-project/package/models"
)

func GetAllWilayahKantor() ([]models.WilayahKantor, error) {
	var ListWilayahKantor []models.WilayahKantor

	err := database.DB.Find(&ListWilayahKantor).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListWilayahKantor, err
}

func GetWilayahKantorById(id string) (models.WilayahKantor, error) {
	var WilayahKantor models.WilayahKantor
	err := database.DB.First(&WilayahKantor, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return WilayahKantor, err
}

func DeleteWilayahKantorById(id string) error {
	err := database.DB.Delete(&models.WilayahKantor{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewWilayahKantor(WilayahKantor models.WilayahKantor) error {
	err := database.DB.Save(&WilayahKantor).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateWilayahKantorById(id string, WilayahKantor models.WilayahKantor) error {
	err := database.DB.Model(&WilayahKantor).Where("id = ?", id).Updates(WilayahKantor).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}