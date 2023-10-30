package usuarios

import (
	"errors"

	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) Save(usuario *entities.Usuario) error {
	err := r.DB.Save(&usuario).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource": "usuarios",
		}

		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return database.NewConflictError(errormessages.CarritoNotBelonging.String())
		}

		return database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}

	return nil
}

func (r Repository) GetByID(usuarioID uint64) (entities.Usuario, error) {
	var usuario entities.Usuario
	err := r.DB.Where("id = ?", usuarioID).Take(&usuario).Error
	if err != nil {
		params := errormessages.Parameters{
			"resource":   "usuarios",
			"usuario_id": usuarioID,
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Usuario{}, database.NewNotFoundError(errormessages.UsuarioNotFound.GetMessageWithParams(params))
		}

		params["error"] = err.Error()
		return entities.Usuario{}, database.NewInternalServerError(errormessages.InternalServerError.GetMessageWithParams(params))
	}

	return usuario, nil
}
