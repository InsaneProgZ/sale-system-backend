package repository

import (
	"sale-system/src/model/domain"
	"time"
)

func Save(product domain.Product) (code int64) {
	sql := `INSERT into products (code, name, buy_value , sell_value, brand, creation_date) values (null, ? , ? , ?, ?, ?);`
	queryResult, err := DB.Exec(
		sql,
		product.Name,
		product.BuyValue,
		product.SellValue,
		product.Brand,
		product.Creation_date)
	if err != nil {
		panic(err.Error())
	}

	code, err = queryResult.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return
}

func FindAll() []domain.Product {
	queryResult, err := DB.Query("SELECT * from products")
	if err != nil {
		panic(err)
	}

	products := []domain.Product{}
	for queryResult.Next() {
		product := domain.Product{}

		queryResult.Scan(&product.Code, &product.Name, &product.BuyValue, &product.SellValue, &product.Brand, &product.Creation_date)
		product.Creation_date = product.Creation_date.In(time.Local)
		products = append(products, product)
		println(product.Code,
			product.Name,
			product.Brand,
			product.BuyValue,
			product.SellValue,
			product.Creation_date.Format("Monday, 02-Jan-06 15:04:05 MST"))
	}
	return products
}
