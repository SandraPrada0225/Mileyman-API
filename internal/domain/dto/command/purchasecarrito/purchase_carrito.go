package purchasecarrito

import (
	"Mileyman-API/internal/domain/dto/contracts/purchasecarrito"
)

type PurchaseCarritoCommand struct {
	CarritoID     uint64
	CompradorID   uint64
	MedioDePagoID uint64
}

func NewPurchaseCarritoCommandFromRequest(request purchasecarrito.Request) PurchaseCarritoCommand {
	return PurchaseCarritoCommand{
		CompradorID:   request.Body.CompradorID,
		CarritoID:     request.URLParams.CarritoID,
		MedioDePagoID: request.Body.MedioDePagoID,
	}
}
