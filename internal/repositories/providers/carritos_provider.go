package providers

import "Mileyman-API/internal/domain/entities"

type CarritosProvider interface {
	GetCarritoByCarritoID(carrito_id uint64) (entities.Carrito, error)
	GetDulceByCarritoIDAndDulceID(carritoID uint64, dulceID uint64) (carritoDulce entities.CarritoDulce, exists bool, err error)
	AddDulceInCarrito(carritoDulce entities.CarritoDulce) (err error)
	DeleteDulceInCarrito(carritoDulce entities.CarritoDulce) (err error)
}
