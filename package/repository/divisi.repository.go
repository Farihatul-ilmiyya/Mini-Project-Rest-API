package repository

import (
	"fmt"
	"mini-project/package/db"
	"mini-project/package/models"
)

func GetAllDivisi() ([]models.Divisi, error) {
	var ListDivisi []models.Divisi

	err := database.DB.Find(&ListDivisi).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListDivisi, err
}

func GetDivisiById(id string) (models.Divisi, error) {
	var Divisi models.Divisi
	err := database.DB.First(&Divisi, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return Divisi, err
}

func DeleteDivisiById(id string) error {
	err := database.DB.Delete(&models.Divisi{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewDivisi(Divisi models.Divisi) error {
	err := database.DB.Save(&Divisi).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateDivisiById(id string, Divisi models.Divisi) error {
	err := database.DB.Model(&Divisi).Where("id = ?", id).Updates(Divisi).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}