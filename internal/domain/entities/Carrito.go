package entities

import estadoscarrito "Mileyman-API/internal/domain/constants/estados_carrito"

type Carrito struct {
	ID              uint64
	Subtotal        float64
	Descuento       float64
	Envio           float64
	PrecioTotal     float64
	EstadoCarritoID uint64
}

func (carrito *Carrito) MarkAsPurchased() {
	carrito.EstadoCarritoID = estadoscarrito.Purchased
}

func NewEmptyCarrito() Carrito {
	return Carrito{
		EstadoCarritoID: estadoscarrito.Active,
	}
}
