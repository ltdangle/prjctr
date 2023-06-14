package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	// 1. Створити словник англійських слів з їх українським перекладом та знайти переклад певного слова.
	dict := make(map[string]string)
	dict["dog"] = "пес"
	dict["cat"] = "кіт"
	dict["table"] = "стіл"
	fmt.Println("Dictionary:")
	for key, value := range dict {
		fmt.Printf("%s - %s\n", key, value)
	}

	// 2. Створити програму, яка буде зберігати дані про працівників певної компанії, додапє і видаляє працівників
	db := newDb()
	e1 := &Employee{FirstName: "John", LastName: "Smith"}
	e2 := &Employee{FirstName: "Michael", LastName: "Jordan"}
	db.addEmployee(e1)
	db.addEmployee(e2)
	db.deleteEmployee(e2.Key)
	fmt.Println("-------")
	fmt.Println("Employee db:")
	// Convert db to json.
	dbJson, _ := json.MarshalIndent(db, "", " ")
	// Print.
	fmt.Println(string(dbJson))
}

// employee Db
type Db struct {
	Employees map[string]*Employee
}

func newDb() *Db {
	db := &Db{}
	db.Employees = make(map[string]*Employee)
	return db
}

func (db *Db) addEmployee(e *Employee) *Employee {
	key := rand.Intn(1000)
	e.Key = strconv.Itoa(key)
	db.Employees[e.Key] = e
	return e
}

func (db *Db) deleteEmployee(key string) {
	delete(db.Employees, key)
}

type Employee struct {
	Key       string `json:"key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
