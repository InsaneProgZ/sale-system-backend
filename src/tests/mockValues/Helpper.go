package mockvalues

import "sale-system/src/model/domain"

func MockDomainProduct() domain.Product {
	return domain.Product{Name: "smart phone", Brand: "Apple", BuyPrice: 1500, SellPrice: 1600}
}
