package models

import (
	"time"
	
)

type Perizinan struct {
	Id_Perizinan        	int       	`gorm:"primaryKey" json:"id_perizinan" form:"id_perizinan"`
	Id_Karyawan        		string    	`gorm:"foreignKey" json:"id_karyawan" form:"id_karyawan"`
	Nama					string		`json:"nama" form:"nama"`
	Tgl_Mulai				time.Time	`json:"tgl_mulai" form:"tgl_mulai"`
	Tgl_Selesai				time.Time	`json:"tgl_selesai" form:"tgl_selesai"`
	Keterangan_Perizinan	string		`json:"keterangan_perizinan" form:"keterangan_perizinan"`
	Created_At         		time.Time 	`json:"created_at" form:"created_at"`
	Updated_At         		time.Time 	`json:"updated_at" form:"updated_at"`
}

func (Perizinan) TableName() string {
	return "Perizinan"
}