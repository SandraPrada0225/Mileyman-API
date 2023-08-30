package dulces

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

func (r Repository) GetByCode(codigo string) (dulce entities.Dulce, err error) {

	err = r.DB.Where("codigo = ?", codigo).First(&dulce).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return entities.Dulce{}, database.NewNotFoundError(errormessages.DulceNotFound)
		}
		return entities.Dulce{},database.NewInternalServerError(errormessages.InternalServerError)
	}

	return 

}
