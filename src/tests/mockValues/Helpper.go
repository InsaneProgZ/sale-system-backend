package mockvalues

import "github.com/InsaneProgZ/sale-system-backend/domain/model"

func MockDomainProduct() model.Product {
	name := "smart phone"
	brand := "Apple"
	Price := uint64(1500)
	return model.Product{Name: name, Brand: brand, Price: Price}
}
