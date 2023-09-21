package query

import "Mileyman-API/internal/domain/entities"

type GetFiltros struct {
	Categorias     []entities.Categoria    `json:"categorias"`
	Marcas         []entities.Marca        `json:"marcas"`
	Presentaciones []entities.Presentacion `json:"presentaciones"`
}
