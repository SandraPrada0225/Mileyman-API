package dulces

import (
	"errors"
	"fmt"

	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

const GetDetalleDulceByCodeSP = "Call GetDetalleDulceByCode(?)"

func (r Repository) GetByCode(codigo string) (detalleDulce query.DetalleDulce, err error) {
	err = r.DB.Raw(GetDetalleDulceByCodeSP, codigo).Take(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource":   "dulces",
			"dulce_code": codigo,
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.DulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInternalServerError(errormessages.DulceNotFound.GetMessageWithParams(params))
		}
	}

	return
}

const GetDetalleDulceByIDSP = "Call GetDetalleDulceByID(?)"

func (r Repository) GetByID(id uint64) (detalleDulce query.DetalleDulce, err error) {
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

func (r Repository) GetDulcesListByCarritoID(carrito_id uint64) ([]uint64, error) {
	var dulcesIDList []uint64

	err := r.DB.Model(&entities.CarritoDulce{}).
		Select("dulce_id").
		Where("carrito_id = ?", carrito_id).
		Find(&dulcesIDList).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource":   "carrito",
			"carrito_id": carrito_id,
			"error":      err.Error(),
		}

		return []uint64{}, database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}

	return dulcesIDList, nil
}
