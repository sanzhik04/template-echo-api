package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"template-echo-api/models"
	"template-echo-api/storage"

	"github.com/labstack/echo/v4"
	"github.com/godruoyi/go-snowflake"
)

	










func generateId() int{
	return int(snowflake.ID())
}

// done
func GetGames(c echo.Context) error {
	var games []models.Game
	err:=storage.DBClient.Select(&games,"SELECT id,name,developer,production FROM games")
	if err!=nil{
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK,games)
} 
// done
func PostGame(c echo.Context) error {
	game:= models.Game{}
	err:= c.Bind(&game)
	if err!= nil{
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	

	game.Id = generateId()
	
	storage.DBClient.Exec("INSERT INTO games (id,name,developer,production) VALUES ($1,$2,$3,$4)", game.Id,game.Name,game.Developer,game.Production)
	if err!= nil{
		c.JSON(http.StatusInternalServerError, err)
	}

	row:=storage.DBClient.QueryRow("SELECT id,name,developer,production FROM games WHERE id = $1",game.Id)
	var AddedGame models.Game
	if err:= row.Scan(&AddedGame.Id,&AddedGame.Name,&AddedGame.Developer,&AddedGame.Production);err!=nil{
		if err == sql.ErrNoRows{
			c.JSON(http.StatusNotFound, err)
		}
		c.JSON(http.StatusUnprocessableEntity, err)
	}



	return c.JSON(http.StatusOK, AddedGame)




}
//done
func GetGameById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var game models.Game
	err:=storage.DBClient.Get(&game,"SELECT id,name,developer,production FROM games WHERE id = $1",id)
	if err!= nil{
		if err == sql.ErrNoRows{
			c.JSON(http.StatusNotFound, err)
		}
		c.JSON(http.StatusUnprocessableEntity, err)
	}
	



	return c.JSON(http.StatusOK, game)
}

//done
func PutGame(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedGame models.Game
	err:=c.Bind(&updatedGame);if err!=nil{
		return c.JSON(http.StatusBadRequest, "Not valid parameters")
	}
	if err:=storage.DBClient.Get(&updatedGame,"UPDATE games SET name = $1, developer = $2, production = $3 WHERE id = $4 RETURNING *",
		updatedGame.Name,
		updatedGame.Developer,
		updatedGame.Production,
		id); err != nil{
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
	
		return c.JSON(http.StatusOK, updatedGame)
}




// done
func DeleteGame(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	
	if _, err := storage.DBClient.Exec("DELETE FROM games WHERE id = $1",id); err!= nil{
		return c.JSON(http.StatusUnprocessableEntity, err)
	}
	return c.JSON(http.StatusOK, id)
}