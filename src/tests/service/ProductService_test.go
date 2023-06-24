package service_test

import (
	"reflect"
	"sale-system/src/model/domain"
	"sale-system/src/service"
	mockvalues "sale-system/src/tests/mockValues"
	"testing"

	"github.com/stretchr/testify/mock"
)

type databaseMock struct {
	mock.Mock
}

func (mock *databaseMock) Save(product domain.Product) int64 {
	args := mock.Called(product)
	return int64(args.Int(0))
}
func (mock *databaseMock) FindAll() []domain.Product {
	return []domain.Product{}
}

func (mock *databaseMock) FindById(id int64) domain.Product {
	return domain.Product{}
}

var dbMock = &databaseMock{}

var productService = &service.ProductServiceImpl{Database: dbMock}

func TestProductServiceImpl_CreateProduct(t *testing.T) {

	product := mockvalues.MockDomainProduct()
	wantProduct := mockvalues.MockDomainProduct()
	wantProduct.Code = 1

	dbMock.On("Save", product).Return(1)

	if got := productService.CreateProduct(product); !reflect.DeepEqual(got, wantProduct) {
		t.Errorf("ProductServiceImpl.CreateProduct() = %v, want %v", got, wantProduct)
	}

}
