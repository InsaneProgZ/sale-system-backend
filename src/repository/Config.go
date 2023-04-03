package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func ConnectSQL() (*sql.DB, error) {
	dbSource := fmt.Sprintf(
		"yan:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		"yan",
		"localhost",
		"3306",
		"sale-system",
	)
	DB, err = sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	return DB, err
}
