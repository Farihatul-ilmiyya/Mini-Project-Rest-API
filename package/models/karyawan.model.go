package models

import (
	"time"
)

type Karyawan struct {
	Id_Karyawan         int       	`gorm:"primaryKey" json:"id_karyawan" form:"id_karyawan"`
	Nama       			string    	`json:"nama" form:"nama"`
	NIK					int			`json:"NIK" form:"nama_divisi"`
	Email				string		`json:"email" form:"email"`
	No_Telepon			int			`json:"no_Telepon" form:"no_Telepon"`
	Alamat				string		`json:"alamat" form:"alamat"`
	Id_Divisi			string		`json:"id_divisi" form:"id_divisi"`
	Id_WilayahKantor	string		`json:"id_wilayahkantor" form:"id_wilayahkantor"`
	Created_At         time.Time `json:"created_at" form:"created_at"`
	Updated_At         time.Time `json:"updated_at" form:"updated_at"`
}

func (Karyawan) TableName() string {
	return "Karyawan"
}