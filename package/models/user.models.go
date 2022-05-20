package models

import (
	"time"
)

type User struct {
	Id_User          int       	`gorm:"primaryKey" json:"id_user" form:"id_user"`
	Username        string    	`json:"username" form:"username"`
	Role_User		string		`json:"role_user" form:"role_user"`
	Password		string		`json:"password" form:"password"`
	Created_At      time.Time 	`json:"created_at" form:"created_at"`
	Updated_At      time.Time 	`json:"updated_at" form:"updated_at"`
}

type UserCostum struct { // table user_custom
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// TableName overrides the table name used by User to `profiles`
func (UserCostum) TableName() string {
	return "users"
}

func (User) TableName() string {
	return "user"
}

