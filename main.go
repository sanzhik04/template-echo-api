package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var games []Game
	





type Game struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Developer string `json:"developer"`
	Production bool `json:"production"`
}

func generateId() int{
	return rand.Int()
}


func getGames(c echo.Context) error {
	return c.JSON(http.StatusOK,games)
} 

func postGame(c echo.Context) error {
	game:= Game{}
	err:= c.Bind(&game)
	if err!= nil{
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	

	game.Id = generateId()
	games = append(games, game)
	
	return c.JSON(http.StatusCreated, games)
}

func getGameById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _,game := range games{
		if game.Id == id {
			return c.JSON(http.StatusOK, game)
		}
	}

	return c.JSON(http.StatusBadRequest, "Game Not Found")
}


func putGame(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	for index,game := range games{
		if  game.Id==id {
			games = append(games[:index],games[index+1:]...)
			var updatedGame Game 
			err:=c.Bind(&updatedGame);if err!=nil{
				return c.JSON(http.StatusBadRequest, "Not valid parameters")
			}
			updatedGame.Id = id
			games = append(games, updatedGame)
			return c.JSON(http.StatusOK, updatedGame)

		}

	}
		return c.JSON(http.StatusBadRequest, "Game Not Found")
}





func deleteGame(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	for index,game := range games{
		if  game.Id== id {
			games = append(games[:index],games[index+1:]...)
			c.JSON(http.StatusOK, games)

		}

	}
		return c.JSON(http.StatusBadRequest, nil)
}



func main(){

	fmt.Println("Running server....")
	e:= echo.New()

	e.GET("/games",getGames)
	e.POST("/games", postGame)
	e.GET("/games/:id",getGameById)
	e.PUT("/games/:id", putGame)
	e.DELETE("/games/:id", deleteGame)
	e.Start(":80")


}