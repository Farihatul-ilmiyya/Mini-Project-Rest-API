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


func (User) TableName() string {
	return "user"
}

//login model
// package models

// import (
// 	"mini-project/package/db"
// 	"fmt"
// 	"mini-project/package/helpers"
// )
// type User struct {
// 	Username        string    	`json:"username" form:"username"`
// 	Password		string		`json:"password" form:"password"`
// }

// func CheckLogin(username, password string)(bool, error){
// 	var obj User
// 	var pwd string

// 	con := db.CreateCon()


// 	sqlStatement := "SELECT * FROM users WHERE username = ?"

// 	err := con.QueryRow(sqlStatement, username).Scan(
// 		&obj.Id, &obj.Username, &pwd,

// 	)

// 	if err == sql.ErrNoRows {
// 		fmt.Println("Username not found")
// 		return false, err	
// 	}

// 	if err != nil{
// 		fmt.Println("Query error")
// 		return false, err
// 	}

// 	match, err := helpers.CheckPasswordHash(password, pwd)
// 		if !match {
// 		fmt.Println("Hash and password doesn't match.")
// 		return false, err
// 	}

// 	return true, nil
// }