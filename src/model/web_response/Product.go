package web_response

import "time"

type Product struct {
	Code         int64     `json:"code"`
	Name         string    `json:"name"`
	BuyValue     uint64    `json:"buy_value"`
	SellValue    uint64    `json:"sell_value"`
	Brand        string    `json:"brand"`
	CreationDate time.Time `json:"creation_date"`
}
