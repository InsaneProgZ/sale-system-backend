package service_test

import (
	"reflect"
	"sale-system/model/domain"
	"sale-system/service"
	mockvalues "sale-system/tests/mockValues"
	"testing"

	"github.com/stretchr/testify/mock"
)

type databaseMock struct {
	mock.Mock
}

func (mock *databaseMock) Save(product domain.Product) (int64, error) {
	args := mock.Called(product)
	code := int64(args.Int(0))
	return code, nil
}
func (mock *databaseMock) FindAll() ([]domain.Product, error) {
	return []domain.Product{}, nil
}

func (mock *databaseMock) FindByCode(id int64) (domain.Product, error) {
	return domain.Product{}, nil
}

func (mock *databaseMock) ChangeProductByCode(id int64, oldProduct domain.Product) (err error) {
	return nil
}

var dbMock = &databaseMock{}

var productService = &service.ProductServiceImpl{Repository: dbMock}

func TestProductServiceImpl_CreateProduct(t *testing.T) {

	product := mockvalues.MockDomainProduct()
	wantProduct := mockvalues.MockDomainProduct()
	code := int64(1)
	wantProduct.Code = code

	dbMock.On("Save", product).Return(1)

	if got, _ := productService.CreateProduct(product); !reflect.DeepEqual(got, wantProduct) {
		t.Errorf("ProductServiceImpl.CreateProduct() = %v, want %v", got, wantProduct)
	}

}
