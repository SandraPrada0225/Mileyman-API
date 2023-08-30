package database;

type InternalServerError struct{
	mensaje  string 
}

func (e InternalServerError) Error()string{
	return e.mensaje
}

func NewInternalServerError (mensaje string) error{
	return InternalServerError{
		mensaje: mensaje,
	}
}