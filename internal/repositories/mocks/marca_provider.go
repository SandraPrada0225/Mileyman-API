package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockMarcaProvider struct {
	mock.Mock
}

func (mock *MockMarcaProvider) GetAll() ([]entities.Marca, error) {
	args := mock.Called()
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.([]entities.Marca), err
	}
	return []entities.Marca{}, err
}
