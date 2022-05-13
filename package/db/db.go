package db
import (
	"database/sql"
	"mini-project/package/config"
	_"github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init(){
	conf := config.GetConfig()

	connectionString := conf.DB_Username + ":" + conf.DB_Password + "@tcp(" + conf.DB_Host + ":" + conf.DB_Port + ")/" + conf.DB_Name
	
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error...")
	}

	err =db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}