package global

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //driver mysql
)

// DB Connecntion
var DB dBConnection

type dBConnection struct {
	Core *sql.DB
}

// InitDB : init db from main
func InitDB() {
	core, err := sql.Open("mysql", "root:[password]@tcp(localhost:3306)/bill?charset=utf8&parseTime=True")
	if err != nil {
		log.Printf("db.user not available, error : %s", err.Error())
		log.Fatal("App exit")
	}
	DB = dBConnection{core}

	_, err = DB.Core.Query("SELECT 1")
	if err != nil {
		log.Printf("Core is not accesible, error: %s", err.Error())
		log.Fatal("App exit")
	}
}
