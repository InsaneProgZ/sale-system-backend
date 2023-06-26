package repository

import (
	"database/sql"
	"sale-system/src/model/domain"
	"time"
)

type Database interface {
	Save(product domain.Product) (int64, error)
	FindAll() ([]domain.Product, error)
	FindByCode(id int64) (domain.Product, error)
	ChangeProductByCode(id int64, product domain.Product) error
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

func (database *MysqlDB) ChangeProductByCode(id int64, newProduct domain.Product) (err error) {

	_, err = database.FindByCode(id)

	if err != nil {
		return
	}

	sql := `UPDATE products SET name = ?, buy_price = ?, sell_price = ?, brand = ? WHERE code = ?`
	_, err = database.Mysql.Exec(
		sql,
		newProduct.Name,
		newProduct.BuyPrice,
		newProduct.SellPrice,
		newProduct.Brand,
		id)
	return
}

// func createUpdateQuery(product domain.Product) (sql string) {
// 	sql = `UPDATE into products SET name = ?, buy_price = ?, sell_price = ?, brand = ? WHERE code = ?;`

// 	for _, field := range product {
// 		switch field {
// 		case "Brand":
// 			sql = strings.Replace(sql, "SET", "SET "+ field + "= ?", 1)
// 		}
// 	}

// }
