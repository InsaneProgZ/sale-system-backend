package model

import (
	"time"
)

type Product struct {
	Code          int64
	Name          string
	Brand         string
	Price         uint64
	Creation_date time.Time
}
