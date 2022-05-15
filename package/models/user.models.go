package models

import (
	"time"
)

type User struct {
	Id_User          int       	`gorm:"primaryKey" json:"id_user" form:"id_user"`
	Nama        	string    	`json:"nama" form:"nama"`
	Role_User		string		`json:"role_user" form:"role_user"`
	Password		string		`json:"password" form:"password"`
	Created_At      time.Time 	`json:"created_at" form:"created_at"`
	Updated_At      time.Time 	`json:"updated_at" form:"updated_at"`
}

func (User) TableName() string {
	return "user"
}