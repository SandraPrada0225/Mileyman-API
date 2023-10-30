package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockCarritoProvider struct {
	mock.Mock
}

func (mock *MockCarritoProvider) GetByID(carrito_id uint64) (entities.Carrito, error) {
	args := mock.Called(carrito_id)
	response := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return entities.Carrito{}, err
	}
	return response.(entities.Carrito), nil
}

func (mock *MockCarritoProvider) Save(carrito *entities.Carrito) error {
	args := mock.Called(carrito)
	createdID := args.Get(0).(uint64)
	if createdID != 0 {
		carrito.ID = createdID
	}
	err := args.Error(1)
	return err
}
