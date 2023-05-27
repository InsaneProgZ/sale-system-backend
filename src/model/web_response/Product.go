package web_response

import "time"

type Product struct {
	Code         int64     `json:"code"`
	Name         string    `json:"name"`
	BuyPrice     uint64    `json:"buy_price"`
	SellPrice    uint64    `json:"sell_price"`
	Brand        string    `json:"brand"`
	CreationDate time.Time `json:"creation_date"`
}
