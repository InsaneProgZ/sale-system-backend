package domain

import "github.com/google/uuid"

type Product struct {
	Name      string
	BuyValue  uint64
	SellValue uint64
	Brand     string
	Code      uuid.UUID
}
