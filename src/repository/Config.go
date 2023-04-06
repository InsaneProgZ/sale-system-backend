package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	dbSource := fmt.Sprintf(
		"yan:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"yan",
		"localhost",
		"3306",
		"sale-system",
	)
	ConnectionDB, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	return ConnectionDB
}
