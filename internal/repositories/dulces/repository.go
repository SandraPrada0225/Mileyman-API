package dulces

import (
	"errors"
	"fmt"

	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/domain/errors/errormessages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

const GetDetalleDulceByCodeSP = "Call GetDetalleDulceByCode(?)"

func (r Repository) GetByCode(codigo string) (detalleDulce responses.DetalleDulce, err error) {
	err = r.DB.Raw(GetDetalleDulceByCodeSP, codigo).Take(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "dulces",
			"dulce_code": codigo,
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}

	return
}

const GetDetalleDulceByIDSP = "Call GetDetalleDulceByID(?)"

func (r Repository) GetDetailByID(id uint64) (detalleDulce responses.DetalleDulce, err error) {
	err = r.DB.Raw(GetDetalleDulceByIDSP, id).Take(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "dulces",
			"dulce_id": fmt.Sprint(id),
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}

	return
}

func (r Repository) GetDulcesListByCarritoID(carrito_id uint64) ([]entities.CarritoDulce, error) {
	var dulcesInCarrito []entities.CarritoDulce

	err := r.DB.Model(&entities.CarritoDulce{}).
		Where("carrito_id = ?", carrito_id).
		Find(&dulcesInCarrito).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource":   "dulces",
			"carrito_id": carrito_id,
			"error":      err.Error(),
		}

		return []entities.CarritoDulce{}, database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}

	return dulcesInCarrito, nil
}

func (r Repository) GetByID(id uint64) (dulce entities.Dulce, err error) {
	err = r.DB.Where("id = ?", id).Take(&dulce).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "dulce",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
		}
	}
	return
}
