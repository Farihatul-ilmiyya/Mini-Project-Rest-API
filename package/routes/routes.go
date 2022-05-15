package routes

import (
	"mini_project/package/controllers"
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Routes Login
	//e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	//e.POST("/createuser", controllers.CreateUserController)

	// Routes divisi
	e.GET("/divisi", controllers.GetAllDivisiController)
	e.GET("/divisi/:id", controllers.GetDivisiController)
	e.DELETE("/divisi/:id", controllers.DeleteDivisiController)
	e.POST("/divisi", controllers.CreateDivisiController)
	e.PUT("/divisi/:id", controllers.UpdateDivisiController)

	// Routes karyawan
	e.GET("/karyawan", controllers.GetAllKaryawanController)
	e.GET("/karyawan/:id", controllers.GetKaryawanController)
	e.DELETE("/karyawan/:id", controllers.DeleteKaryawanController)
	e.POST("/karyawan", controllers.CreateKaryawanController)
	e.PUT("/karyawan/:id", controllers.UpdateKaryawanController)

	// Routes Penjadwalan
	e.GET("/penjadwalan", controllers.GetAllPenjadwalanController)
	e.GET("/penjadwalan/:id", controllers.GetPenjadwalanController)
	e.DELETE("/penjadwalan/:id", controllers.DeletePenjadwalanController)
	e.POST("/penjadwalan", controllers.CreatePenjadwalanController)
	e.PUT("/penjadwalan/:id", controllers.UpdatePenjadwalanController)

	// Routes Perizinan
	e.GET("/perizinan", controllers.GetAllPerizinanController)
	e.GET("/perizinan/:id", controllers.GetPerizinanController)
	e.DELETE("/perizinan/:id", controllers.DeletePerizinanController)
	e.POST("/perizinan", controllers.CreatePerizinanController)
	e.PUT("/perizinan/:id", controllers.UpdatePerizinanController)

	//Routes Presensi
	e.GET("/presensi", controllers.GetAllPresensiController)
	e.GET("/presensi/:id", controllers.GetPresensiController)
	e.DELETE("/presensi/:id", controllers.DeletePresensiController)
	e.POST("/presensi", controllers.CreatePresensiController)
	e.PUT("/presensi/:id", controllers.UpdatePresensiController)

	// Routes User
	e.GET("/user", controllers.GetAllUserController)
	e.GET("/user/:id", controllers.GetUserController)
	e.DELETE("/user/:id", controllers.DeleteUserController)
	e.POST("/user", controllers.CreateUserController)
	e.PUT("/user/:id", controllers.UpdateUserController)

	// Routes Wilayah Kantor
	e.GET("/wilayahkantor", controllers.GetAllWilayahKantorController)
	e.GET("/wilayahkantor/:id", controllers.GetTipeWilayahKantorController)
	e.DELETE("/wilayahkantor/:id", controllers.DeleteWilayahKantorController)
	e.POST("/wilayahkantor", controllers.CreateWilayahKantorController)
	e.PUT("/wilayahkantor/:id", controllers.UpdateWilayahKantorController)
	
	return e
}