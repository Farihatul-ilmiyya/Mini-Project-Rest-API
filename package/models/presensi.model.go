package models

import (
	"time"
)

type Presensi struct {
	Id_Presensi         int       	`gorm:"primaryKey" json:"id_presensi" form:"id_presensi"`
	Id_Karyawan        	string    	`gorm:"foreignKey" json:"id_karyawan" form:"id_karyawan"`
	Tgl_Presensi		time.Time	`json:"tgl_presensi" form:"tgl_presensi"`
	Jam_Masuk			time.Time	`json:"jam_masuk" form:"jam_masuk"`
	Jam_Keluar			time.Time	`json:"jam_keluar" form:"jam_keluar"`
	Status_Masuk		string		`json:"status_masuk" form:"status_masuk"`
	Status_Keluar		string		`json:"status_keluar" form:"status_keluar"`
	Waktu_Telat			time.Time	`json:"waktu_telat" form:"waktu_telat"`
	Status_Kehadiran	string		`json:"status_kehadiran" form:"status_kehadiran"`
	Created_At         	time.Time 	`json:"created_at" form:"created_at"`
	Updated_At         	time.Time 	`json:"updated_at" form:"updated_at"`
}

func (Presensi) TableName() string {
	return "presensi"
}