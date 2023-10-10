package errormessages

import "fmt"

type (
	ErrorMessage string
	Parameters   map[string]interface{}
)

const (
	CarritoDulceNotFound    ErrorMessage = "No se encontró un detalle carrito_dulce con ese codigo"
	DulceNotFound           ErrorMessage = "No se encontró un dulce con ese codigo"
	InternalServerError     ErrorMessage = "Ha ocurrido un error inesperado"
	IdMustBeAPositiveNumber ErrorMessage = "el ID debe ser un número positivo"
	UnitLimitExceded        ErrorMessage = "las unidades requeridad exceden las disponibles"
	CarritoNotFound     ErrorMessage = "No se encontró un carrito con este id"
	InvalidTypeError    ErrorMessage = "El tipo de dato es inválido"
)

func (e ErrorMessage) String() string {
	return string(e)
}

func (e ErrorMessage) GetMessageWithParams(params Parameters) string {
	msg := e.String()

	for key, value := range params {
		msg = fmt.Sprintf("%s. %s: %s", msg, key, value)
	}

	return msg
}
