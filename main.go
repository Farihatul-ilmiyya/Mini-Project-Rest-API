package main

import (
	"mini-project/package/db"
	"mini-project/package/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}