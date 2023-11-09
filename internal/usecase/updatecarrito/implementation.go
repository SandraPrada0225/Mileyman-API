package updatecarrito

import (
	"Mileyman-API/internal/domain/constants"
	"Mileyman-API/internal/domain/dto/command/updatecarrito"
	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/business"
	"Mileyman-API/internal/domain/errors/errormessages"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	CarritosProvider providers.CarritosProvider
	DulcesProvider   providers.DulcesProvider
}

func (UseCase Implementation) Execute(carritoID uint64, movements updatecarrito.Body) (responses.MovementsResult, error) {
	_, err := UseCase.CarritosProvider.GetByID(carritoID)
	if err != nil {
		return responses.MovementsResult{}, err
	}
	movementsResult := responses.NewMovementsResult()
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
		case !exists:
			operationResult = constants.Created
			carritoDulce = entities.NewCarritoDulce(carritoID, movement.DulceID)

			err := UseCase.save(movement, carritoDulce)
			if err != nil {
				movementsResult.AddResult(index, movement.DulceID, constants.Error.String(), err.Error())
				continue
			}
		case exists:
			operationResult = constants.Updated
			err := UseCase.save(movement, carritoDulce)
			if err != nil {
				movementsResult.AddResult(index, movement.DulceID, constants.Error.String(), err.Error())
				continue
			}
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
