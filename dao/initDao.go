package dao

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test")
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	mysqlDB = db
	err = mysqlDB.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}

}
