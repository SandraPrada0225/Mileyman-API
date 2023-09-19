package providers

import "Mileyman-API/internal/domain/entities"

type MarcasProvider interface {
	GetAll() (marca []entities.Marca, err error)
}
