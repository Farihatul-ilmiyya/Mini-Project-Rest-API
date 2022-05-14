package models

import (
	"time"
)

type Divisi struct {
	Id_Karyawan         int       `gorm:"primaryKey" json:"id_divisi" form:"id_divisi"`
	Nama       			string    `json:"nama_divisi" form:"nama_divisi"`
	NIK					int
	Email				string
	No_Telepon
	Alamat
	Id_Divisi
	Id_WilayahKantor
	Created_At         time.Time `json:"created_at" form:"created_at"`
	Updated_At         time.Time `json:"updated_at" form:"updated_at"`
}

func (Divisi) TableName() string {
	return "Divisi"
}