package repository

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB
var err error

func Config() {
	DB, err = sql.Open("mysql", "root:<yourMySQLdatabasepassword>@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer DB.Close()
	fmt.Println("Success!")
}
