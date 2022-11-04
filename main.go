package main

import (
	"fmt"
	

	"template-echo-api/storage"
	"template-echo-api/controllers"

	"github.com/labstack/echo/v4"
)





func main(){

	
	fmt.Println("Running server....")
	e:= echo.New()
	storage.NewDB()


	
	
	e.GET("/games", controllers.GetGames)
	e.POST("/games", controllers.PostGame)
	e.GET("/games/:id", controllers.GetGameById)
	e.PUT("/games/:id",  controllers.PutGame)
	e.DELETE("/games/:id",  controllers.DeleteGame)
	e.Start(":4000")


}