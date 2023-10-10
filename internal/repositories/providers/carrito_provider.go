package providers

import "Mileyman-API/internal/domain/entities"

type CarritoProvider interface {
	GetCarritoByCarritoID(carrito_id uint64) (entities.Carrito, error)
}
