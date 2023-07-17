package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	port := ":8080"

	classA := NewClass()
	seedClass(1, classA, 30)
	classB := NewClass()
	seedClass(2, classB, 30)

	school := NewSchool()
	school.AddClass(classA)
	school.AddClass(classB)

	rspndr := NewResponder("2006-01-02 15:04:05")
	app := NewApp(school, rspndr)

	http.Handle("/", app.auth(http.HandlerFunc(app.httpHandler)))

	fmt.Println("Server started on port " + port)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func seedClass(teacherId int, class *Class, numStudents int) {
	// Create teacher.
	teacher := NewTeacher(teacherId, "Teacher_"+strconv.Itoa(teacherId))
	teacher.Username = "Teacher_" + strconv.Itoa(teacherId)
	teacher.Password = "Teacher_" + strconv.Itoa(teacherId)

	// Create students.
	for i := 0; i < numStudents; i++ {
		studentId := rand.Intn(1000)
		s := NewStudent(studentId, fmt.Sprintf("Student%d", studentId))
		class.AddStudent(s)
	}

	class.SetTeacher(teacher)
}
