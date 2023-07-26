package dulces

import (
	"Mileyman-API/internal/domain/entities"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetByCode(codigo string) (dulce entities.Dulce, err error) {

	err = r.DB.Where("codigo = ?", codigo).First(&dulce).Error

	return

}
