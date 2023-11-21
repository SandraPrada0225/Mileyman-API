package responses

import (
	"time"
)

type GetPurchaseList struct {
	PurchaseList []Purchase `json:"purchase_list"`
}

type Purchase struct {
	ID              uint64    `json:"id"`
	CarritoID       uint64    `json:"carrito_id"`
	Fecha           time.Time `json:"fecha"`
	MedioDePagoID   uint64    `json:"medio_de_pago_id"`
	MedioDePago     string    `json:"medio_de_pago"`
	PrecioTotal     float64   `json:"precio_total"`
	Subtotal        float64   `json:"sobtotal"`
	Descuento       float64   `json:"descuento"`
	Envio           float64   `json:"envio"`
	EstadoCarritoID uint64    `json:"estado_carrito_id"`
	EstadoCarrito   string    `json:"estado_carrito"`
}
