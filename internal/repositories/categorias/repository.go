package categorias

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/domain/errors/errormessages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetAll() (categorias []entities.Categoria, err error) {
	err = r.DB.Find(&categorias).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "categorias",
			"error":    err.Error(),
		}
		return []entities.Categoria{}, database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}

	return
}

const GetCategoriasByDulceIDSP = "Call GetCategoriasByDulceID(?)"

func (r Repository) GetCategoriasByDulceID(dulceID uint64) (categorias []entities.Categoria, err error) {
	err = r.DB.Raw(GetCategoriasByDulceIDSP, dulceID).Scan(&categorias).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "categorias",
		}
		err = database.NewInternalServerError(errormessages.DulceNotFound.GetMessageWithParams(params))
	}

	return
}
