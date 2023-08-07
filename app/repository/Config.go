package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(dbUrl string) *sql.DB {
	dbSource := fmt.Sprintf(
		"yan:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"yan",
		dbUrl,
		"3306",
		"sale-system",
	)
	ConnectionDB, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	return ConnectionDB
}
