package categorias

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (categorias []entities.Categoria, err error) {
	err = r.DB.Find(&categorias).Error
	if err != nil {
		return []entities.Categoria{}, database.NewInternalServerError(errormessages.InternalServerError)
	}

	return
}
