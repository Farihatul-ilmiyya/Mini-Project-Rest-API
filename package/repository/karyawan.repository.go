package repository

import (
	"fmt"
	"mini-project/package/db"
	"mini-project/package/models"
)

func GetAllKaryawan() ([]models.Karyawan, error) {
	var ListKaryawan []models.Karyawan

	err := database.DB.Find(&ListKaryawan).Error
	if err != nil {
		fmt.Println(err)
	}
	return ListKaryawan, err
}

func GetKaryawanById(id string) (models.Karyawan, error) {
	var Karyawan models.Karyawan
	err := database.DB.First(&Karyawan, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return Karyawan, err
}

func DeleteKaryawanById(id string) error {
	err := database.DB.Delete(&models.Karyawan{}, "id = ?", id).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func CreateNewKaryawan(Karyawan models.Karyawan) error {
	err := database.DB.Save(&Karyawan).Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func UpdateKaryawanById(id string, Karyawan models.Karyawan) error {
	err := database.DB.Model(&Karyawan).Where("id = ?", id).Updates(Karyawan).Debug().Error
	if err != nil {
		fmt.Println(err)
	}
	return err
}