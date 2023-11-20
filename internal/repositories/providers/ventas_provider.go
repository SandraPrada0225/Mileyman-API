package providers

import (
	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/entities"
)

type VentasProvider interface {
	Create(venta *entities.Venta) error
	GetListByUserID(userID uint64) (responses.GetPurchaseList, error)
}
