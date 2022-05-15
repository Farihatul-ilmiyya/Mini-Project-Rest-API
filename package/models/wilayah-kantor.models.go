package models

import (
	"time"
)

type WilayahKantor struct {
	Id_WilayahKantor     	int 	 	`gorm:"primaryKey" json:"id_wilayahkantor" form:"id_wilayahkantor"`
	Nama_Wilayah        	string   	`json:"nama_wilayah" form:"nama_wilayah"`
	Created_At         		time.Time 	`json:"created_at" form:"created_at"`
	Updated_At         		time.Time 	`json:"updated_at" form:"updated_at"`
}

func (WilayahKantor) TableName() string {
	return "WilayahKantor"
}