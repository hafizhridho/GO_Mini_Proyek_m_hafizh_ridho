package main

import (
	"latihan/configs"
	"latihan/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	//configs.Loadenv()
	configs.InitDb()
	e := echo.New()
	e.POST("/users", controllers.Register)
	e.POST("/login", controllers.LoginController)

	eAUTH := e.Group("")
	
	eAUTH.Use(echojwt.JWT([]byte("secretKey")))
	eAUTH.POST("/list", controllers.CreateList)
	eAUTH.GET("/list", controllers.GetAllLists)
	eAUTH.PUT("/list/:id", controllers.UpdateList)
	eAUTH.DELETE("/list/:id", controllers.DeleteList)
	eAUTH.GET("/list/:id", controllers.GetListByID)
	

	eAUTH.POST("/tugas", controllers.CreateTugas)
	eAUTH.GET("/tugas", controllers.GetTugas)
	eAUTH.GET("/tugas/:id", controllers.GetTaskById)
	eAUTH.PUT("/tugas/:id", controllers.UpdateTask)
	eAUTH.PUT("/tugas/status/:id", controllers.UpdateTugasStatus)
	

	e.Start(":8000")
	
}