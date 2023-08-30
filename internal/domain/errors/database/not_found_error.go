package database;

type NotFoundError struct{
	mensaje  string 
}

func (e NotFoundError) Error()string{
	return e.mensaje
}

func NewNotFoundError (mensaje string) error{
	return NotFoundError{
		mensaje: mensaje,
	}
}