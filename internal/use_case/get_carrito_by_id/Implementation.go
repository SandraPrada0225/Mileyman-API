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
	carrito, err := useCase.CarritoProvider.GetByID(id)
	if err != nil {
		return query.GetDetalleCarrito{}, err
	}

	dulcesIDList, err := useCase.DulcesProvider.GetDulcesListByCarritoID(id)
	if err != nil {
		return query.GetDetalleCarrito{}, err
	}

	var dulcesList []query.DulceInCarrito

	for _, dulceInCarrito := range dulcesIDList {
		dulce, err := useCase.DulcesProvider.GetDetailByID(dulceInCarrito.DulceID)
		if err != nil {
			return query.GetDetalleCarrito{}, err
		}

		categoriasList, err := useCase.CategoriasProvider.GetCategoriasByDulceID(dulce.ID)
		if err != nil {
			return query.GetDetalleCarrito{}, err
		}

		dulce.AddCategorias(categoriasList)
		dulceInCarritoItem := query.NewDulceInCarrito(dulce, dulceInCarrito.Unidades, dulceInCarrito.Subtotal)

		dulcesList = append(dulcesList, dulceInCarritoItem)
	}

	return query.NewGetDetalleCarrito(carrito, dulcesList), nil
}
