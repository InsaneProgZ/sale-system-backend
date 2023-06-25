package repository

import (
	"database/sql"
	"log"
	"sale-system/src/model/domain"
	"time"
)

type Database interface {
	Save(product domain.Product) (int64, error)
	FindAll() ([]domain.Product, error)
	FindById(id int64) (domain.Product, error)
}

type MysqlDB struct {
	Mysql *sql.DB
}

func (database *MysqlDB) Save(product domain.Product) (code int64, err error) {
	sql := `INSERT into products (code, name, buy_price , sell_price, brand, creation_date) values (null, ? , ? , ?, ?, ?);`
	queryResult, err := database.Mysql.Exec(
		sql,
		product.Name,
		product.BuyPrice,
		product.SellPrice,
		product.Brand,
		product.Creation_date)
	if err != nil {
		log.Println(err)
		return
	}

	code, err = queryResult.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return
}

func (database *MysqlDB) FindAll() (products []domain.Product, err error) {

	queryResult, err := database.Mysql.Query("SELECT * from products")
	if err != nil {
		log.Println(err)
		return
	}

	for queryResult.Next() {
		product := domain.Product{}

		queryResult.Scan(&product.Code, &product.Name, &product.BuyPrice, &product.SellPrice, &product.Brand, &product.Creation_date)
		product.Creation_date = product.Creation_date.In(time.Local)
		products = append(products, product)
		println(product.Code,
			product.Name,
			product.Brand,
			product.BuyPrice,
			product.SellPrice,
			product.Creation_date.Format("Monday, 02-Jan-06 15:04:05 MST"))
	}
	return
}

func (database *MysqlDB) FindById(id int64) (product domain.Product, err error) {

	queryResult := database.Mysql.QueryRow("SELECT * from products where code=?", id)

	err = queryResult.Scan(&product.Code, &product.Name, &product.BuyPrice, &product.SellPrice, &product.Brand, &product.Creation_date)

	product.Creation_date = product.Creation_date.In(time.Local)

	return
}
