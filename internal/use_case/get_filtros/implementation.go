package getfiltros

import (
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	CategoriasProvider     providers.CategoriasProvider
	MarcasProvider         providers.MarcasProvider
	PresentacionesProvider providers.PresentacionProvider
}

func (UseCase Implementation) Execute() (query.GetFiltros, error) {
	marcas, err := UseCase.MarcasProvider.GetAll()
	if err != nil {
		return query.GetFiltros{}, err
	}
	presentaciones, err := UseCase.PresentacionesProvider.GetAll()
	if err != nil {
		return query.GetFiltros{}, err
	}
	categorias, err := UseCase.CategoriasProvider.GetAll()
	if err != nil {
		return query.GetFiltros{}, err
	}

	return query.GetFiltros{
		Marcas:         marcas,
		Presentaciones: presentaciones,
		Categorias:     categorias,
	}, nil
}
