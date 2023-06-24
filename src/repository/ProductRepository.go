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
	Mysql sql.DB
}

func (db *MysqlDB) Save(product domain.Product) (code int64, err error) {
	ConnectionDB := ConnectDB()
	sql := `INSERT into products (code, name, buy_price , sell_price, brand, creation_date) values (null, ? , ? , ?, ?, ?);`
	queryResult, err := ConnectionDB.Exec(
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
	defer ConnectionDB.Close()

	code, err = queryResult.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return
}

func (db *MysqlDB) FindAll() ([]domain.Product, error) {
	ConnectionDB := ConnectDB()
	queryResult, err := ConnectionDB.Query("SELECT * from products")
	if err != nil {
		panic(err)
	}
	defer ConnectionDB.Close()

	products := []domain.Product{}
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
	return products, nil
}

func (db *MysqlDB) FindById(id int64) (domain.Product, error) {
	ConnectionDB := ConnectDB()
	queryResult := ConnectionDB.QueryRow("SELECT * from products where code=?", id)

	defer ConnectionDB.Close()

	product := domain.Product{}
	queryResult.Scan(&product.Code, &product.Name, &product.BuyPrice, &product.SellPrice, &product.Brand, &product.Creation_date)
	product.Creation_date = product.Creation_date.In(time.Local)
	// println(product.Code, product.Name, product.BuyPrice, product.SellPrice, product.Brand, product.Creation_date.String())
	return product, nil
}
