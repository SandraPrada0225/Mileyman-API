package marcas

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (marcas []entities.Marca, err error) {
	err = r.DB.Find(&marcas).Error
	if err != nil {
		return []entities.Marca{}, database.NewInternalServerError(errormessages.InternalServerError)
	}
	return
}