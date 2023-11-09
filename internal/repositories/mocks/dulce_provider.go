package mocks

import (
	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockDulceProvider struct {
	mock.Mock
}

func (mock *MockDulceProvider) GetByCode(codigo string) (responses.DetalleDulce, error) {
	args := mock.Called(codigo)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(responses.DetalleDulce), err
	}
	return responses.DetalleDulce{}, err
}

func (mock *MockDulceProvider) GetDetailByID(id uint64) (responses.DetalleDulce, error) {
	args := mock.Called(id)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(responses.DetalleDulce), err
	}
	return responses.DetalleDulce{}, err
}

func (mock *MockDulceProvider) GetDulcesListByCarritoID(carrito_id uint64) ([]entities.CarritoDulce, error) {
	args := mock.Called(carrito_id)
	response := args.Get(0)
	err := args.Error(1)
	if response != nil {
		return response.([]entities.CarritoDulce), err
	}
	return []entities.CarritoDulce{}, err
}

func (mock *MockDulceProvider) GetByID(id uint64) (entities.Dulce, error) {
	args := mock.Called(id)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(entities.Dulce), err
	}
	return entities.Dulce{}, err
}
