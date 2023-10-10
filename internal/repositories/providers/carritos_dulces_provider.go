package providers

import "Mileyman-API/internal/domain/entities"

type CarritosDulcesProvider interface {
	GetDulceByCarritoIDAndDulceID(carritoID uint64, dulceID uint64) (carritoDulce entities.CarritoDulce, exists bool, err error)
	AddDulceInCarrito(carritoDulce entities.CarritoDulce) (err error)
	DeleteDulceInCarrito(carritoDulce entities.CarritoDulce) (err error)
}
