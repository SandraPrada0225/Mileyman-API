package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockUsuarioProvider struct {
	mock.Mock
}

func (m *MockUsuarioProvider) Save(usuario *entities.Usuario) error {
	args := m.Called(usuario)
	err := args.Error(0)
	return err
}

func (m *MockUsuarioProvider) GetByID(usuarioID uint64) (entities.Usuario, error) {
	args := m.Called(usuarioID)
	response := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return entities.Usuario{}, err
	}

	return response.(entities.Usuario), nil
}
