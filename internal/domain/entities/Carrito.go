package entities

type Carrito struct {
	ID          uint64
	SubTotal    float64
	Descuento   float64
	Envio       float64
	PrecioTotal float64
}
