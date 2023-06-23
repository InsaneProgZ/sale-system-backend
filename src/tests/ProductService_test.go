package service

import (
	"reflect"
	"sale-system/src/model/domain"
	"sale-system/src/service"
	"testing"
)

type databaseMock struct{}

func (mock *databaseMock) Save(product domain.Product) int64 {
	return 1
}
func (mock *databaseMock) FindAll() []domain.Product {
	return []domain.Product{}
}

func (mock *databaseMock) FindById(id int64) domain.Product {
	return domain.Product{}
}

func TestProductServiceImpl_CreateProduct(t *testing.T) {
	type args struct {
		product domain.Product
	}
	tests := []struct {
		name           string
		productService *service.ProductServiceImpl
		args           args
		want           any
	}{
		{
			name:           "Saved with success",
			productService: &service.ProductServiceImpl{Database: &databaseMock{}},
			args:           args{domain.Product{Name: "smart phone", Brand: "Apple", BuyPrice: 1500, SellPrice: 1600}},
			want:           domain.Product{Code: 1, Name: "smart phone", Brand: "Apple", BuyPrice: 1500, SellPrice: 1600},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.productService.CreateProduct(tt.args.product); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductServiceImpl.CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
