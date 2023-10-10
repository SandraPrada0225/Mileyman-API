package getcarritobyid

import (
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	CarritoProvider    providers.CarritoProvider
	DulcesProvider     providers.DulcesProvider
	CategoriasProvider providers.CategoriasProvider
}

func (useCase Implementation) Execute(id uint64) (query.GetDetalleCarrito, error) {
	carrito, err := useCase.CarritoProvider.GetCarritoByCarritoID(id)
	if err != nil {
		return query.GetDetalleCarrito{}, err
	}

	dulcesIDList, err := useCase.DulcesProvider.GetDulcesListByCarritoID(id)
	if err != nil {
		return query.GetDetalleCarrito{}, err
	}

	var dulcesList []query.DetalleDulce

	for _, dulceID := range dulcesIDList {
		dulce, err := useCase.DulcesProvider.GetDetailByID(dulceID)
		if err != nil {
			return query.GetDetalleCarrito{}, err
		}

		categoriasList, err := useCase.CategoriasProvider.GetCategoriasByDulceID(dulce.ID)
		if err != nil {
			return query.GetDetalleCarrito{}, err
		}

		dulce.AddCategorias(categoriasList)
		dulcesList = append(dulcesList, dulce)
	}

	return query.NewGetDetalleCarrito(carrito, dulcesList), nil
}
