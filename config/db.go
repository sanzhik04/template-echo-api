package config

import (
	"fmt"

)


const (
	username = "root"
	password = "root"
	database = "root"
	host = "0.0.0.0"
	port = 5432
)

func GetPostgresConnectionString() string {
	connectionString:= fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",host,port,username,password,database)
	return connectionString

}