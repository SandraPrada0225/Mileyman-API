package getcarritobyid

import (
	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	CarritoProvider    providers.CarritosProvider
	DulcesProvider     providers.DulcesProvider
	CategoriasProvider providers.CategoriasProvider
}

func (useCase Implementation) Execute(id uint64) (responses.GetDetalleCarrito, error) {
	carrito, err := useCase.CarritoProvider.GetByID(id)
	if err != nil {
		return responses.GetDetalleCarrito{}, err
	}

	dulcesIDList, err := useCase.DulcesProvider.GetDulcesListByCarritoID(id)
	if err != nil {
		return responses.GetDetalleCarrito{}, err
	}

	var dulcesList []responses.DulceInCarrito

	for _, dulceInCarrito := range dulcesIDList {
		dulce, err := useCase.DulcesProvider.GetDetailByID(dulceInCarrito.DulceID)
		if err != nil {
			return responses.GetDetalleCarrito{}, err
		}

		categoriasList, err := useCase.CategoriasProvider.GetCategoriasByDulceID(dulce.ID)
		if err != nil {
			return responses.GetDetalleCarrito{}, err
		}

		dulce.AddCategorias(categoriasList)
		dulceInCarritoItem := responses.NewDulceInCarrito(dulce, dulceInCarrito.Unidades, dulceInCarrito.Subtotal)

		dulcesList = append(dulcesList, dulceInCarritoItem)
	}

	return responses.NewGetDetalleCarrito(carrito, dulcesList), nil
}
