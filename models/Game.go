package models

type Game struct{
	Id int `json:"id" db:"id"` 
	Name string `json:"name" db:"name"`
	Developer string `json:"developer" db:"developer"`
	Production bool `json:"production" db:"production"`
}



