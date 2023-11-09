package presentaciones

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/domain/errors/errormessages"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (presentaciones []entities.Presentacion, err error) {
	err = r.DB.Find(&presentaciones).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "marcas",
			"error":    err.Error(),
		}
		return []entities.Presentacion{}, database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}
	return
}
