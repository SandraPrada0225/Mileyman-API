package entities

import "time"

type Venta struct {
	ID            uint64
	MedioDePagoID uint64
	CarritoID     uint64
	CompradorID   uint64
	CreatedAt     time.Time
}

func (Venta) TableName() string {
	return "ventas"
}

func NewVenta(medioDePago, carritoID, compradorID uint64) Venta {
	return Venta{
		MedioDePagoID: medioDePago,
		CarritoID:     carritoID,
		CompradorID:   compradorID,
	}
}
