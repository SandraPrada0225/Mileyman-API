package business

type CarritoAlreadyPurchaseError struct {
	mensaje string
}

func (e CarritoAlreadyPurchaseError) Error() string {
	return e.mensaje
}

func NewCarritoAlreadyPurchasedError(mensaje string) error {
	return CarritoAlreadyPurchaseError{
		mensaje: mensaje,
	}
}
