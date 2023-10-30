package providers

import "Mileyman-API/internal/domain/entities"

type UsuariosProvider interface {
	Save(usuario *entities.Usuario) error
	GetByID(usuarioID uint64) (entities.Usuario, error)
}
