package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockCategoriaProvider struct {
	mock.Mock // implementacion boba de la interface//herencia de otra forma en go
}

func (mock *MockCategoriaProvider) GetAll() ([]entities.Categoria, error) {
	args := mock.Called()
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.([]entities.Categoria), err
	}
	return []entities.Categoria{}, err
}
