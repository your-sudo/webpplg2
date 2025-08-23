package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DBconnection() {
	db, err := sql.Open("mysql", "root:root@/web_kelas")
	if err != nil {
		panic(err.Error())
	}

	log.Println("Database connected")
	DB = db
	
}