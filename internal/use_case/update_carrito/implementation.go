package updatecarrito

import (
	"Mileyman-API/internal/domain/constants"
	updatecarrito "Mileyman-API/internal/domain/dto/command/update_carrito"
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/business"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	CarritosProvider providers.CarritosProvider
	DulcesProvider   providers.DulcesProvider
}

func (UseCase Implementation) Execute(carritoID uint64, movements updatecarrito.Body) (query.MovementsResult, error) {
	_, err := UseCase.CarritosProvider.GetCarritoByCarritoID(carritoID)
	if err != nil {
		return query.MovementsResult{}, err
	}
	movementsResult := query.NewMovementsResult()
	for index, movement := range movements.Movements {
		var operationResult constants.CarritoOperationResult
		carritoDulce, exists, err := UseCase.CarritosProvider.GetDulceByCarritoIDAndDulceID(carritoID, movement.DulceID)
		if err != nil {
			movementsResult.AddResult(index, movement.DulceID, constants.Error.String(), err.Error())
			continue
		}
		switch {
		case movement.Unidades == 0:
			operationResult = constants.Deleted
			err = UseCase.CarritosProvider.DeleteDulceInCarrito(carritoDulce)
			if err != nil {
				movementsResult.AddResult(index, movement.DulceID, constants.Error.String(), err.Error())
				continue
			}
			break
		case !exists:
			operationResult = constants.Created
			carritoDulce = entities.NewCarritoDulce(carritoID, movement.DulceID)

			err := UseCase.save(movement, carritoDulce)
			if err != nil {
				movementsResult.AddResult(index, movement.DulceID, constants.Error.String(), err.Error())
				continue
			}
			break
		case exists:
			operationResult = constants.Updated
			err := UseCase.save(movement, carritoDulce)
			if err != nil {
				movementsResult.AddResult(index, movement.DulceID, constants.Error.String(), err.Error())
				continue
			}
			break
		}
		movementsResult.AddResult(index, movement.DulceID, operationResult.String(), "")
	}
	return movementsResult, err
}

func (UseCase Implementation) save(movement updatecarrito.Movement, carritoDulce entities.CarritoDulce) error {
	dulce, err := UseCase.DulcesProvider.GetByID(movement.DulceID)
	if err != nil {
		return err
	}

	if movement.Unidades > dulce.Disponibles {
		return business.NewUnitLimitExceded(errormessages.UnitLimitExceded.String())
	}

	carritoDulce = entities.UpdateCarritoDulce(carritoDulce, movement.Unidades, dulce.PrecioUnidad)

	err = UseCase.CarritosProvider.AddDulceInCarrito(carritoDulce)
	return err
}