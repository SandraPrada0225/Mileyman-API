package providers

import "Mileyman-API/internal/domain/entities"

type CarritoProvider interface {
	GetByID(carrito_id uint64) (entities.Carrito, error)
	Save(carrito *entities.Carrito) error
}
