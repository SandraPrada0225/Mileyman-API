package errormessages

import "fmt"

type (
	ErrorMessage string
	Parameters   map[string]interface{}
)

const (
	DulceNotFound       ErrorMessage = "No se encontró un dulce con ese codigo"
	InternalServerError ErrorMessage = "Ha currido un error inesperado"
)

func (e ErrorMessage) String() string {
	return string(e)
}

func (e ErrorMessage) GetMessageWithParams(params Parameters) string {
	msg := e.String()

	for key, value := range params {
		msg = fmt.Sprintf("%s.\n%s - %s", msg, key, value)
	}

	return msg
}
