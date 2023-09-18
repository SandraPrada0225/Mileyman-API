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

const GetDetalleDulceByCodeSP = "Call GetCategoriasByDulceID(?)"

func (r *Repository) GetCategoriasByDulceID(dulceID uint64) (categorias []entities.Categoria, err error) {
	err = r.DB.Raw(GetDetalleDulceByCodeSP, dulceID).Scan(&categorias).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "dulces",
		}
		err = database.NewInternalServerError(errormessages.DulceNotFound.GetMessageWithParams(params))
	}

	return
}
