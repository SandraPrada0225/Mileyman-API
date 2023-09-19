package providers

import "Mileyman-API/internal/domain/entities"

type CategoriasProvider interface {
	GetAll() (categoria []entities.Categoria, err error)
}
