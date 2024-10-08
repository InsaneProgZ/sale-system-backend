package mysql

import (
	"database/sql"
	"log"
	"reflect"
	"time"

	"github.com/InsaneProgZ/sale-system-backend/domain/model"
)

type Database interface {
	Save(product model.Product) (int64, error)
	FindAll() ([]model.Product, error)
	FindByCode(code int64) (model.Product, error)
	ChangeProductByCode(code int64, product model.Product) error
}

type MysqlDB struct {
	Mysql *sql.DB
}

func (database *MysqlDB) Save(product model.Product) (code int64, err error) {
	log.Default()
	sql := `INSERT into products (code, name, price, brand, creation_date) values (null, ? , ?, ?, ?);`
	queryResult, err := database.Mysql.Exec(
		sql,
		product.Name,
		product.Price,
		product.Brand,
		product.Creation_date)
	if err != nil {
		return
	}
	code, err = queryResult.LastInsertId()
	return
}

func (database *MysqlDB) FindAll() (products []model.Product, err error) {
	queryResult, err := database.Mysql.Query("SELECT * from products")
	if err != nil {
		return
	}

	for queryResult.Next() {
		product := model.Product{}

		queryResult.Scan(&product.Code, &product.Name, &product.Price, &product.Brand, &product.Creation_date)
		localTime := product.Creation_date.In(time.Local)
		product.Creation_date = localTime
		products = append(products, product)
	}
	return
}

func (database *MysqlDB) FindByCode(id int64) (product model.Product, err error) {
	queryResult := database.Mysql.QueryRow("SELECT * from products where code=?", id)

	err = queryResult.Scan(&product.Code, &product.Name, &product.Price, &product.Brand, &product.Creation_date)
	if err != nil {
		return
	}
	localTime := product.Creation_date.In(time.Local)
	product.Creation_date = localTime
	return
}

func (database *MysqlDB) ChangeProductByCode(code int64, newProduct model.Product) (err error) {

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

func createUpdateQuery(code int64, product model.Product) (sql string, params []interface{}) {
	sql = `UPDATE products SET`

	productType := reflect.TypeOf(product)
	productValue := reflect.ValueOf(product)

	for i := 0; i < productType.NumField(); i++ {
		switch productType.Field(i).Name {
		case "Name":
			if !productValue.Field(i).IsNil() {
				sql += " name = ?,"
				params = append(params, *productValue.Field(i).Interface().(*string))
			}
		case "Price":
			if !productValue.Field(i).IsNil() {
				sql += " price = ?,"
				params = append(params, *productValue.Field(i).Interface().(*uint64))
			}
		case "Brand":
			if !productValue.Field(i).IsNil() {
				sql += " brand = ?,"
				params = append(params, *productValue.Field(i).Interface().(*string))
			}
		default:

		}
	}
	sql = sql[:len(sql)-1]
	sql += " WHERE code = ?"
	params = append(params, code)
	return
}
