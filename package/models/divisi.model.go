package models

import (
	"time"
	"mini-project/package/db"
)

type divisi struct {
	id_divisi 		int       `gorm:"primaryKey" json:"id" form:"id"`
	nama_divisi 	string    `json:"nama_divisi" form:"nama_divisi"`
	Created_At      time.Time `json:"created_at" form:"created_at"`
	Updated_At      time.Time `json:"updated_at" form:"updated_at"`
}

func FetchAllDivisi()(Response, error){
	var obj divisi
	var arrobj []divisi
	var res Response

	con :=db.CreateCon()
	sqlStatement := "SELECT * From divisi"
}