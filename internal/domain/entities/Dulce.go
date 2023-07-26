package entities

import (
	"time"
)

type Dulce struct {
	ID               uint64
	Nombre           string
	MarcaID          uint64
	Precio           int
	Peso             float64
	Unidades         int
	PresentacionID   uint64
	Descripcion      string
	Imagen           string
	FechaVencimiento time.Time
	FechaExpedicion  time.Time
	Disponibles      int
	codigo           string
}


