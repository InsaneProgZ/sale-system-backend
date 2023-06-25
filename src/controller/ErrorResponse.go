package controller

type ErrorResponse struct {
	Field   string
	Message string
}

var ValidationsMessage = map[string]string{
	"required": "required",
	"gt":       "greater than",
}
