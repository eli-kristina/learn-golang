package main

import (
	"database/sql"
	"fmt"
)

import _ "github.com/go-sql-driver/mysql"

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

func main()  {
	db, err := connect()
	
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, name, grade FROM students WHERE age = ?", 18)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []student

	for rows.Next() {
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.id, each.name, each.grade)
	}
}