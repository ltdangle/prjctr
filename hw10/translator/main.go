package main

import (
	"fmt"
	"log"
	"net/http"
)

type Translation struct {
	FromLng     string
	ToLng       string
	Source      string
	Translation string
}

func main() {
	_, tr := translate("en", "es", "good evening")
	fmt.Println(tr)
	// var v interface{}
	// err := json.Unmarshal([]byte(jsonData), &v)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	return
	port := ":8080"

	rspndr := NewResponder("2006-01-02 15:04:05")
	app := NewApp(rspndr)

	http.HandleFunc("/translate", app.httpHandler)

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
