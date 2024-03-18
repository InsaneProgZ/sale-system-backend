package web_response

import "time"

type Product struct {
	Code         int64     `json:"code"`
	Name         string    `json:"name"`
	Price        uint64    `json:"price"`
	Brand        string    `json:"brand"`
	CreationDate time.Time `json:"creation_date"`
}
