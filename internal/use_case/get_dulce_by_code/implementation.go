package getdulcebycode

import (
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	DulcesProvider     providers.DulcesProvider
	CategoriasProvider providers.CategoriasProvider
}

func (UseCase Implementation) Execute(codigo string) (query.DetalleDulce, error) {

	response, err := UseCase.DulcesProvider.GetByCode(codigo)
	if err != nil {
		return query.DetalleDulce{}, err
	}
	categorias, err := UseCase.CategoriasProvider.GetCategoriasByDulceID(response.ID)
	if err != nil {
		return query.DetalleDulce{}, err
	}

	response.AddCategorias(categorias)

	return response, nil
}
