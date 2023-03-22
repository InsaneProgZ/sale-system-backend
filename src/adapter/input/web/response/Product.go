package web

import "github.com/google/uuid"

type Product struct {
	Code  uuid.UUID `json:"code"`
	Name  string    `json:"name"`
	Value uint64    `json:"value"`
	Brand string    `json:"brand"`
}
	