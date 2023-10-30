package command

import (
	"Mileyman-API/cmd/server/handlers/purchase_carrito/contract"
)

type PurchaseCarritoCommand struct {
	CarritoID     uint64
	CompradorID   uint64
	MedioDePagoID uint64
}

func NewPurchaseCarritoCommandFromRequest(request contract.Request) PurchaseCarritoCommand {
	return PurchaseCarritoCommand{
		CompradorID:   request.Body.CompradorID,
		CarritoID:     request.URLParams.CarritoID,
		MedioDePagoID: request.Body.MedioDePagoID,
	}
}
