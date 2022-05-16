package repository

import (
	"fmt"
	"mini-project/package/db"
	"mini-project/package/models"
)

func GetAllPresensi() ([]models.Presensi, error) {
	var ListPresensi []models.Presensi

	err := database.DB.Find(&ListPresensi).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListPresensi, err
}

func GetPresensiById(id string) (models.Presensi, error) {
	var Presensi models.Presensi
	err := database.DB.First(&Presensi, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return Presensi, err
}

func DeletePresensiById(id string) error {
	err := database.DB.Delete(&models.Presensi{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewPresensi(Presensi models.Presensi) error {
	err := database.DB.Save(&Presensi).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdatePresensiById(id string, Presensi models.Presensi) error {
	err := database.DB.Model(&Presensi).Where("id = ?", id).Updates(Presensi).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}