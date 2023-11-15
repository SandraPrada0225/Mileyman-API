package marcas

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/domain/errors/errormessages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (marcas []entities.Marca, err error) {
	err = r.DB.Find(&marcas).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "marcas",
			"error":    err.Error(),
		}
		return []entities.Marca{}, database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return
}
