package dulces

import (
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

const GetDetalleDulceByCodeSP = "Call GetDetalleDulceByCode(?)"

func (r Repository) GetByCode(codigo string) (detalleDulce query.DetalleDulce, err error) {

	err = r.DB.Raw(GetDetalleDulceByCodeSP, codigo).Scan(&detalleDulce).Error

	if err != nil {
		params := errormessages.Parameters{
			"resource": "dulces",
		}
		err = database.NewInternalServerError(errormessages.DulceNotFound.GetMessageWithParams(params))
	}

	return
}
