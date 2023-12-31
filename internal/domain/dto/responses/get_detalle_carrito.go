package responses

import "Mileyman-API/internal/domain/entities"

type GetDetalleCarrito struct {
	ID          uint64           `json:"id"`
	Subtotal    float64          `json:"subtotal"`
	Descuento   float64          `json:"descuento"`
	Envio       float64          `json:"envio"`
	PrecioTotal float64          `json:"precio_total"`
	DulcesList  []DulceInCarrito `json:"dulces_list"`
}

func NewGetDetalleCarrito(carrito entities.Carrito, dulcesList []DulceInCarrito) GetDetalleCarrito {
	return GetDetalleCarrito{
		ID:          carrito.ID,
		Subtotal:    carrito.Subtotal,
		Descuento:   carrito.Descuento,
		Envio:       carrito.Envio,
		PrecioTotal: carrito.PrecioTotal,
		DulcesList:  dulcesList,
	}
}
