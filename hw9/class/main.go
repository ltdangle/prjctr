package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Username string
	Password string
}

var teacherUser = User{
	Username: "teacher",
	Password: "teacher",
}

func main() {
	port := ":8080"

	class := NewClass()
	seedClass(class, 30)

	// Dump data in console for debugging
	classJsonData, _ := json.MarshalIndent(class, "", " ")
	fmt.Println(string(classJsonData))

	rspndr := NewResponder("2006-01-02 15:04:05")
	app := NewApp(class, rspndr)

	http.Handle("/", auth(http.HandlerFunc(app.httpHandler)))

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func seedClass(class *Class, numStudents int) {
	for i := 0; i < numStudents; i++ {
		s := NewStudent(i, fmt.Sprintf("Student%d", i))

		class.addStudent(s)
	}
}


