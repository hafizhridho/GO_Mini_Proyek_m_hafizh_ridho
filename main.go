package main

import (
	"latihan/configs"
	"latihan/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func main() {
	configs.Loadenv()
	configs.InitDb()
	e := echo.New()
	e.POST("/users", controllers.Register)
	e.POST("/login", controllers.Login)

	eAUTH := e.Group("")
	eAUTH.Use(echojwt.JWT([]byte("")))
	eAUTH.POST("/list", controllers.CreateList)
	e.GET("/list", controllers.GetAllLists)
	e.PUT("/list/:id", controllers.UpdateList)
	e.DELETE("/list/:id", controllers.DeleteList)
	e.GET("/list/:id", controllers.GetListByID)

	e.POST("/tugas", controllers.CreateTugas)
	e.GET("/tugas", controllers.GetTugas)
	e.PUT("/tugas/:id", controllers.UpdateTask)
	e.PUT("/tugas/status/:id", controllers.UpdateTugasStatus)

	e.Start(":8000")
	
}