package database

type ConflictError struct {
	mensaje string
}

func (e ConflictError) Error() string {
	return e.mensaje
}

func NewConflictError(mensaje string) error {
	return ConflictError{
		mensaje: mensaje,
	}
}
