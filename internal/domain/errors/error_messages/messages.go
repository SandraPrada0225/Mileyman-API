package errormessages

import "fmt"

type (
	ErrorMessage string
	Parameters   map[string]interface{}
)

const (
	DulceNotFound                  ErrorMessage = "No se encontró un dulce con ese codigo"
	CarritoNotFound                ErrorMessage = "No se encontró un carrito con este id"
	UsuarioNotFound                ErrorMessage = "No se encontró un usuario con este id"
	CarritoNotBelonging            ErrorMessage = "El carrito pertenece a otro usuario"
	InternalServerError            ErrorMessage = "Ha currido un error inesperado"
	InvalidTypeError               ErrorMessage = "El tipo de dato es inválido"
	CarritoHasAlreadyBeenPurchased ErrorMessage = "El carrito ya ha sido comprado"
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
