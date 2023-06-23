package service

import (
	"reflect"
	"sale-system/src/model/domain"
	"testing"
)

func TestProductServiceImpl_CreateProduct(t *testing.T) {
	type args struct {
		product domain.Product
	}
	tests := []struct {
		name           string
		productService *ProductServiceImpl
		args           args
		want           domain.Product
	}{
		{
			name: "S",
			productService: &ProductServiceImpl{}, //mock it
			args: args{domain.Product{Code: 1, Name: "d12", Brand: "21", BuyPrice: 1,SellPrice: 2}},
			want: domain.Product{Code: 1, Name: "d12", Brand: "21", BuyPrice: 1,SellPrice: 2},
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
