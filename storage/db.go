package storage

import (
	"fmt"
	config "template-echo-api/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


var DBClient *sqlx.DB

func NewDB(params... string){
	db, err:= sqlx.Connect("postgres", config.GetPostgresConnectionString())
	if err != nil{
		fmt.Println("Connection to DB Error: ", err)

	}

	//defer db.Close()

	//db.SetConnMaxLifetime(30 * time.Second)
	//db.SetConnMaxIdleTime(30)

	DBClient = db

}