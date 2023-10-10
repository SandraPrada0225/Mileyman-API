package carritos

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// consultamos la existencia del dulce carrito_dulce
func (r Repository) GetDulceByCarritoIDAndDulceID(carritoID uint64, dulceID uint64) (carritoDulce entities.CarritoDulce, exists bool, err error) {
	err = r.DB.Where("carrito_id = ? AND dulce_id = ?", carritoID, dulceID).First(&carritoDulce).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "carrito",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exists = false
			err = nil
		} else {
			err = database.NewInternalServerError(errormessages.CarritoDulceNotFound.GetMessageWithParams(params))
		}
		return
	}
	exists = true;
	return
}

// agregamos dulce al carrito o lo editamos
func (r Repository) AddDulceInCarrito(carritoDulce entities.CarritoDulce) (err error) {
	err = r.DB.Save(&carritoDulce).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "carrito",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.CarritoDulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInternalServerError(errormessages.CarritoDulceNotFound.GetMessageWithParams(params))
		}
	}
	return
}

// eliminamos dulce del carrito
func (r Repository) DeleteDulceInCarrito(carritoDulce entities.CarritoDulce) (err error) {
	err = r.DB.Delete(&carritoDulce).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "carrito",
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = database.NewNotFoundError(errormessages.CarritoDulceNotFound.GetMessageWithParams(params))
		} else {
			err = database.NewInternalServerError(errormessages.CarritoDulceNotFound.GetMessageWithParams(params))
		}
	}
	return
}
