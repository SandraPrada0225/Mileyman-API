package getfiltros

import (
	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	CategoriasProvider     providers.CategoriasProvider
	MarcasProvider         providers.MarcasProvider
	PresentacionesProvider providers.PresentacionProvider
}

func (UseCase Implementation) Execute() (responses.GetFiltros, error) {
	marcas, err := UseCase.MarcasProvider.GetAll()
	if err != nil {
		return responses.GetFiltros{}, err
	}
	presentaciones, err := UseCase.PresentacionesProvider.GetAll()
	if err != nil {
		return responses.GetFiltros{}, err
	}
	categorias, err := UseCase.CategoriasProvider.GetAll()
	if err != nil {
		return responses.GetFiltros{}, err
	}

	return responses.GetFiltros{
		Marcas:         marcas,
		Presentaciones: presentaciones,
		Categorias:     categorias,
	}, nil
}
