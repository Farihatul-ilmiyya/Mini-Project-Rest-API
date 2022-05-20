package main

import (
	"mini-project/package/db"
	"mini-project/package/routes"
	"mini-project/package/repository"
	"mini-project/package/controllers"
	"mini-project/package/middleware"
)

func init() {
	database.InitDBConnect()
	database.InitialMigration()
}

func main() {
	iUserRepo := repository.NewUserRepo(database.DB) // --> layer data   <-- unit test

	loginController := controllers.NewLoginController(iUserRepo) // --> layer controller
	userController := controllers.NewUserController(iUserRepo)   // --> layer controller
	e := routes.Init()
	e.POST("/login", loginController.Login)

	e.GET("/users", userController.GetAllData, middleware.ValidateJwt())
	e.GET("/user", userController.GetSingleData, middleware.ValidateJwt())
	e.POST("/users", userController.Create, middleware.ValidateJwt())
	
	

	e.Logger.Fatal(e.Start(":2625"))
}

