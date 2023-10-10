package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockCarritoProvider struct {
	mock.Mock
}

func (mock *MockCarritoProvider) GetCarritoByCarritoID(carrito_id uint64) (entities.Carrito, error) {
	args := mock.Called(carrito_id)
	response := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return entities.Carrito{}, err
	}
	return response.(entities.Carrito), nil
}
