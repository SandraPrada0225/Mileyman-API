package ventas

import (
	"errors"

	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/domain/errors/errormessages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) Create(venta *entities.Venta) error {
	err := r.DB.Create(venta).Error
	if err != nil {

		params := errormessages.Parameters{
			"resource":   "ventas",
			"carrito_id": venta.CarritoID,
		}

		if errors.Is(gorm.ErrForeignKeyViolated, err) {
			return database.NewNotFoundError(errormessages.CarritoNotFound.GetMessageWithParams(params))
		}

		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return database.NewConflictError(errormessages.CarritoHasAlreadyBeenPurchased.GetMessageWithParams(params))
		}

		params["error"] = err.Error()
		return database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}

	return nil
}
