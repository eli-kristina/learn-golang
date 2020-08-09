package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type categories struct {
	id  int
	name string
	description string
	created_at string
	created_by string
	updated_at string
	updated_by string
}

func main () {
	sqlQuery()
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bidding")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = categories{}
	var id = 2

	err = db.
		QueryRow("Select id, name, description from categories where id = ?", id).
		Scan(&result.id, &result.name, &result.description)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("id: %d\nname: %s\ndescription: %s", result.id, result.name, result.description)
}