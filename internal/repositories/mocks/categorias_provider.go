package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockCategoriasProvider struct {
	mock.Mock
}

func (m *MockCategoriasProvider) GetCategoriasByDulceID(dulceID uint64) ([]entities.Categoria, error) {
	responseArgs := m.Called(dulceID)
	response := responseArgs.Get(0)
	err := responseArgs.Error(1)
	return response.([]entities.Categoria), err
}
