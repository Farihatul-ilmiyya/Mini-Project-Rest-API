package models

import (
	"time"
)

type Penjadwalan struct {
	Id_Jadwal        int       `gorm:"primaryKey" json:"id_jadwal" form:"id_jadwal"`
	Jam_Masuk        time.Time  `json:"jam_masuk" form:"jam_masuk"`
	Jam_Keluar	     time.Time	`json:"jam_keluar" form:"jam_keluar"`
	Created_At       time.Time `json:"created_at" form:"created_at"`
	Updated_At       time.Time `json:"updated_at" form:"updated_at"`
}

func (Penjadwalan) TableName() string {
	return "Penjadwalan"
}