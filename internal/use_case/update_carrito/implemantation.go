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
	CarritosDulcesProvider providers.CarritosDulcesProvider
	DulcesProvider         providers.DulcesProvider
}

func (UseCase Implementation) Execute(carritoID uint64, movements updatecarrito.Body) query.MovementsResult {
	movementsResult := query.NewMovementsResult()
	for index, movement := range movements.Movements {
		var operationResult constants.CarritoOperationResult
		carritoDulce, exists, err := UseCase.CarritosDulcesProvider.GetDulceByCarritoIDAndDulceID(carritoID, movement.DulceID)
		if err != nil {
			movementsResult.AddResult(index, movement.DulceID, constants.Error.String(), err.Error())
			continue
		}
		switch {
		case movement.Unidades == 0:
			operationResult = constants.Deleted
			err = UseCase.CarritosDulcesProvider.DeleteDulceInCarrito(carritoDulce)
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
	return movementsResult
}

func (UseCase Implementation) save(movement updatecarrito.Movement, carritoDulce entities.CarritoDulce) error {
	dulce, err := UseCase.DulcesProvider.GetByID(movement.DulceID)
	if err != nil {
		return err
	}

	if movement.Unidades > dulce.Unidades {
		return business.NewUnitLimitExceded(errormessages.UnitLimitExceded.String())
	}

	carritoDulce = entities.UpdateCarritoDulce(carritoDulce, movement.Unidades, dulce.Precio)

	err = UseCase.CarritosDulcesProvider.AddDulceInCarrito(carritoDulce)
	return err
}
