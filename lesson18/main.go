package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//prjctr
	// root:root
	config := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "prjctr",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	log.Default().Println("Pinged!")

}
