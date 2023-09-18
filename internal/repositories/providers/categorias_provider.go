package providers

import "Mileyman-API/internal/domain/entities"

type CategoriasProvider interface {
	GetCategoriasByDulceID(dulceID uint64) (categorias []entities.Categoria, err error)
}
