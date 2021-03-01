package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"
)

import _ "github.com/go-sql-driver/mysql"

type M map[string]interface{}

type student struct {
	id int
	name string
	age int
	grade int
}

var DB_HOST = "127.0.0.1"
var DB_PORT = "3306"
var DB_USER = "root"
var DB_PASS = "root"
var DB_NAME = "learn_golang"

func connect() (*sql.DB, error)  {
	db, err := sql.Open("mysql", DB_USER + ":" + DB_PASS + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func jsonResponse(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write([]byte("\n"))
}

func getStudent(id string) M  {
	db, err := connect()
	
	if err != nil {
		fmt.Println(err.Error())
	}

	var result = student{}

	err = db.
		QueryRow("SELECT id, name, age, grade FROM students WHERE id = ?", id).
		Scan(&result.id, &result.name, &result.age, &result.grade)

	if err != nil {
		fmt.Println(err.Error())
	}

	return M{"id": result.id, "name": result.name, "age": result.age, "grade": result.grade}
}

func main()  {
	http.HandleFunc("/api/students", func(w http.ResponseWriter, r *http.Request) {
		if id := r.URL.Query().Get("id"); id != "" {
			jsonResponse(w, getStudent(id))
		} else {
			http.Error(w, "Invalid id", http.StatusBadRequest)
		}
	})

	fmt.Println("server started at localhost:9090")
	http.ListenAndServe(":9090", nil)
}