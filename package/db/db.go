package database

import (
	"mini-project/package/config"
	"mini-project/package/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDBConnect() {
	conf := config.GetConfig()

	dsn := conf.DB_Username + ":" + conf.DB_Password + "@tcp(" + conf.DB_Host + ":" + conf.DB_Port + ")/" + conf.DB_Name + "?parseTime=true"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func DBConnect() *gorm.DB {
	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&models.Divisi{})
}