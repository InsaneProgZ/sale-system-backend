package response

type BadRequestResponse struct {
	Field   string
	Message string
}

type GenericErrorResponse struct {
	Message string
}

var ValidationsMessage = map[string]string{
	"required": "required",
	"gt":       "greater than",
}

var CodeErrors = map[string]string{
	"name_exists":   "1062",
	"id_not_exists": "no rows in result",
}
