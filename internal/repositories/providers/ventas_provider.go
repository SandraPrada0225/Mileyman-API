package providers

import "Mileyman-API/internal/domain/entities"

type VentasProvider interface {
	Create(venta *entities.Venta) error
}
