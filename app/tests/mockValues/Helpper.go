package mockvalues

import "sale-system/model/domain"

func MockDomainProduct() domain.Product {
	name := "smart phone"
	brand := "Apple"
	Price := uint64(1500)
	return domain.Product{Name: name, Brand: brand, Price: Price}
}
