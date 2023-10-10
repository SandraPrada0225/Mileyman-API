package rest

type BadRequestError struct {
	mensaje string
}

func (e BadRequestError) Error() string {
	return e.mensaje
}

func NewBadRequestError(mensaje string) error {
	return BadRequestError{
		mensaje: mensaje,
	}
}
