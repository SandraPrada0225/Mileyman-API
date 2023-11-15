package entities

type Usuario struct {
	ID              uint64
	Nombre          string
	Apellido        string
	Password        string
	Correo          string
	CarritoActualID uint64
}

func (u *Usuario) UpdateCarritoID(carritoID uint64) {
	u.CarritoActualID = carritoID
}
