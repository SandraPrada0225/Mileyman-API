package getdulcebycode

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	DulcesProvider providers.DulcesProvider
}

func (UseCase Implementation) Execute(codigo string) (entities.Dulce, error) {
	return UseCase.DulcesProvider.GetByCode(codigo)
}
