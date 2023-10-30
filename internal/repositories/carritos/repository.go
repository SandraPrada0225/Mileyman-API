package carritos

import (
	"errors"
	"fmt"

	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetByID(carrito_id uint64) (entities.Carrito, error) {
	var carrito entities.Carrito

	err := r.DB.Where("id = ?", carrito_id).Take(&carrito).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource":   "carrito",
			"carrito_id": fmt.Sprint(carrito_id),
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Carrito{}, database.NewNotFoundError(errormessages.CarritoNotFound.GetMessageWithParams(params))
		}

		params["error"] = err.Error()
		return entities.Carrito{}, database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}

	return carrito, nil
}

func (r Repository) Save(carrito *entities.Carrito) error {
	err := r.DB.Save(carrito).Error
	if err != nil {
		err = database.NewInternalServerError(errormessages.InternalServerError.String())
	}

	return err
}
