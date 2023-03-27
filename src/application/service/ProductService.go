package service

import (
	"sale-system/src/adapter/output/repository"
	"sale-system/src/application/domain"
	"sale-system/src/application/handler"

	"github.com/google/uuid"
)

type IProductService interface {
	CreateProduct(product domain.Product) *domain.Product
}

type productService struct {
	ProductRepository repository.IRespository
}

func NewProductService(productRepository repository.IRespository) IProductService {
	return &productService{
		ProductRepository: productRepository,
	}
}

func (p *productService) CreateProduct(product domain.Product) *domain.Product {
	product.Code, _ = uuid.NewUUID()
	
	err := p.ProductRepository.Save(product)
	handler.ErrorHandler(err)

	return &product
}
