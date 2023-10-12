package entities

import (
	"time"
)

type Dulce struct {
	ID               uint64
	Nombre           string
	MarcaID          uint64
	PrecioUnidad     float64
	Peso             float64
	Unidades         int
	PresentacionID   uint64
	Descripcion      string
	Imagen           string
	FechaVencimiento time.Time
	FechaExpedicion  time.Time
	Disponibles      int
	Codigo           string
}
