package repository

import (
	"database/sql"
	"sale-system/src/application/domain"
)

type IRespository interface {
	Save(product domain.Product) error
}

type productRepository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) IRespository {
	return &productRepository{db}
}

func (r *productRepository) Save(product domain.Product) error {
	_, err := DB.Query("SELECT *")
	return err
}
