package consts

type Header string

var Headers = map[Header]string{
	"Content-Type":                 "application/json",
	"Access-Control-Allow-Origin":  "*",
	"Access-Control-Allow-Headers": "*",
	"Access-Control-Allow-Methods": "*",
}


const (
	ContentType               Header = "Content-Type"
	AccessControlAllowOrigin  Header = "Access-Control-Allow-Origin"
	AccessControlAllowHeaders Header = "Access-Control-Allow-Headers"
	AccessControlAllowMethods Header = "Access-Control-Allow-Methods"
)

func GetHeader(Header Header) (key string, value string) {
	return string(Header), Headers[Header]
}