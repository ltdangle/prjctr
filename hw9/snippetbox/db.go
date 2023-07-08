package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/snippetbox?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	defer db.Close()

	// Run query.
	rows, err := db.Query("select id, title from snippets")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Read query data.
	var (
		id    int
		title string
	)
	for rows.Next() {
		err := rows.Scan(&id, &title)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Insert data.
	stmt, err := db.Prepare("INSERT INTO snippets(title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), UTC_TIMESTAMP())")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Dolly-"+time.Now().String(),"Dolly content")
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}
