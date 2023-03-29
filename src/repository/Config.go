package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB = new(sql.DB)

func ConnectSQL() (*sql.DB, error) {
	dbSource := fmt.Sprintf(
		"yan:%s@tcp(%s:%s)/%s?charset=utf8",
		"yan",
		"localhost",
		"3306",
		"sale-system",
	)
	d, err := sql.Open("mysql", dbSource)
	if err != nil {
		panic(err)
	}
	DB = d
	return DB, err
}
