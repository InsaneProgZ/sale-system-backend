package mockvalues

import "sale-system/src/model/domain"

func MockDomainProduct() domain.Product {
	name := "smart phone"
	brand := "Apple"
	buyPrice := uint64(1500)
	sellPrice := uint64(1600)
	return domain.Product{Name: &name, Brand: &brand, BuyPrice: &buyPrice, SellPrice: &sellPrice}
}
