package repository

import (
	"sale-system/src/model/domain"
)

func Save(product domain.Product) error {
	_, err := DB.Query("SELECT *")
	return err
}