package main

import (
	"mini-project/package/db"
	"mini-project/package/routes"
)

func init() {
	database.InitDBConnect()
	database.InitialMigration()
}

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":2625"))
}

