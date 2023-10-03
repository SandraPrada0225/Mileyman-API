package providers

// es una interface que me provee un conjunto de funciones sin una implementacion especifica

import (
	"Mileyman-API/internal/domain/dto/query"
)

// la interface es un contrato
type DulcesProvider interface {
	GetByCode(codigo string) (dulce query.DetalleDulce, err error)
	GetByID(id uint64) (dulce query.DetalleDulce, err error)
	GetDulcesListByCarritoID(carrito_id uint64) ([]uint64, error)
}
