package repository

import (
	"fmt"
	"sale-system/src/model/domain"
)

func Save(product domain.Product) error {
	_, err := DB.Query("SELECT *")
	return err
}

func FindAll(){
	all, err := DB.Query("SELECT * from products")
	if(err != nil){
		panic(err)
	}
	var name, value string
	all.Next()
	all.Scan(&name, &value)
	fmt.Println(name, value)
}