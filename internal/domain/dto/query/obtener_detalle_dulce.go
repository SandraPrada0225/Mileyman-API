package query

import "Mileyman-API/internal/domain/entities"

type DetalleDulce struct {
	ID               uint64                `json:"id"`
	Peso             int                   `json:"peso"`
	PrecioUnidad     int                   `json:"precio_unidad"`
	Disponibles      int                   `json:"disponibles"`
	Subtotal         int                   `json:"subtotal"`
	Codigo           string                `json:"codigo"`
	Nombre           string                `json:"nombre"`
	Descripcion      string                `json:"descripcion"`
	Imagen           string                `json:"imagen"`
	FechaVencimiento string                `json:"fecha_vencimiento"`
	FechaExpedicion  string                `json:"fecha_expedicion"`
	Categorias       []entities.Categoria  `json:"categorias" gorm:"-"`
	Presentacion     entities.Presentacion `json:"presentacion" gorm:"embedded;embeddedPrefix:presentacion_"`
	Marca            entities.Marca        `json:"marca" gorm:"embedded;embeddedPrefix:marca_"`
}

func(query *DetalleDulce) AddCategorias(categorias []entities.Categoria) {
	query.Categorias = categorias
}