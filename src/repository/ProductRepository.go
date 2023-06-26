package repository

import (
	"database/sql"
	"log"
	"reflect"
	"sale-system/src/model/domain"
	"time"
)

type Database interface {
	Save(product domain.Product) (int64, error)
	FindAll() ([]domain.Product, error)
	FindByCode(code int64) (domain.Product, error)
	ChangeProductByCode(code int64, product domain.Product) error
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
		return
	}

	code, err = queryResult.LastInsertId()
	return
}

func (database *MysqlDB) FindAll() (products []domain.Product, err error) {
	queryResult, err := database.Mysql.Query("SELECT * from products")
	if err != nil {
		return
	}

	for queryResult.Next() {
		product := domain.Product{}

		queryResult.Scan(&product.Code, &product.Name, &product.BuyPrice, &product.SellPrice, &product.Brand, &product.Creation_date)
		product.Creation_date = product.Creation_date.In(time.Local)
		products = append(products, product)
	}
	return
}

func (database *MysqlDB) FindByCode(id int64) (product domain.Product, err error) {
	queryResult := database.Mysql.QueryRow("SELECT * from products where code=?", id)

	err = queryResult.Scan(&product.Code, &product.Name, &product.BuyPrice, &product.SellPrice, &product.Brand, &product.Creation_date)
	product.Creation_date = product.Creation_date.In(time.Local)
	return
}

func (database *MysqlDB) ChangeProductByCode(code int64, newProduct domain.Product) (err error) {

	_, err = database.FindByCode(code)

	if err != nil {
		return
	}

	sql, params := createUpdateQuery(code, newProduct)

	_, err = database.Mysql.Exec(
		sql,
		params...)
	return
}

func createUpdateQuery(code int64, product domain.Product) (sql string, params []interface{}) {
	sql = `UPDATE products SET`

	productType := reflect.TypeOf(product)
	productValue := reflect.ValueOf(product)

	for i := 0; i < productType.NumField(); i++ {
		log.Println(sql)
		switch productType.Field(i).Name {
		case "Name":
			sql += " name = ?,"
			params = append(params, productValue.Field(i).Interface())
		case "BuyPrice":
			sql += " buy_price = ?,"
			params = append(params, productValue.Field(i).Interface())
		case "SellPrice":
			sql += " sell_price = ?,"
			params = append(params, productValue.Field(i).Interface())
		case "Brand":
			sql += " brand = ?,"
			params = append(params, productValue.Field(i).Interface())
		}
	}
	sql = sql[:len(sql)-1]
	sql += " WHERE code = ?"
	params = append(params, code)
	return
}
