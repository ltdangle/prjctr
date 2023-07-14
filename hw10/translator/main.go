package main

import (
	"fmt"
	"log"
	"net/http"
)

type Translation struct {
	FromLng     string `json:"from"`
	ToLng       string `json:"to"`
	Source      string `json:"source"`
	Translation string `json:"translation"`
}

func main() {
	port := ":8080"

	rspndr := NewResponder("2006-01-02 15:04:05")
	trnsltr := NewTranslator()
	app := NewApp(trnsltr, rspndr)

	http.HandleFunc("/translate", app.httpHandler)

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(port, nil))

}
