package business

type UnitLimitExceded struct {
	mensaje string
}

func (e UnitLimitExceded) Error() string {
	return e.mensaje
}

func NewUnitLimitExceded(mensaje string) error {
	return UnitLimitExceded{
		mensaje: mensaje,
	}
}
