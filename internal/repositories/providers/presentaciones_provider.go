package providers

import "Mileyman-API/internal/domain/entities"

type PresentacionProvider interface {
	GetAll() (presentacion []entities.Presentacion, err error)
}
