package providers
//es una interface que me provee un conjunto de funciones sin una implementacion especifica 

import "Mileyman-API/internal/domain/entities"

//la interface es un contrato
type DulcesProvider interface {
	GetByCode(codigo string) (dulce entities.Dulce, err error)
}