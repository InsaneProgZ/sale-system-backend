package response

import "github.com/google/uuid"

type Product struct {
	Code      uuid.UUID `json:"code"`
	Name      string    `json:"name"`
	BuyValue  uint64    `json:"buy_value"`
	SellValue uint64    `json:"sell_value"`
	Brand     string    `json:"brand"`
}
