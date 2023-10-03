package mocks

// dado unos parametros devuelve la respuesta deseada para comprobar que el caso de uso reaccione bien nate las diferentes situaciones

import (
	"Mileyman-API/internal/domain/dto/query"

	"github.com/stretchr/testify/mock"
)

type MockDulceProvider struct {
	mock.Mock // implementacion boba de la interface
}

func (mock *MockDulceProvider) GetByCode(codigo string) (query.DetalleDulce, error) {
	args := mock.Called(codigo)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(query.DetalleDulce), err
	}
	return query.DetalleDulce{}, err
}

func (mock *MockDulceProvider) GetByID(id uint64) (query.DetalleDulce, error) {
	args := mock.Called(id)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(query.DetalleDulce), err
	}
	return query.DetalleDulce{}, err
}

func (mock *MockDulceProvider) GetDulcesListByCarritoID(carrito_id uint64) ([]uint64, error) {
	args := mock.Called(carrito_id)
	response := args.Get(0)
	err := args.Error(1)
	if response != nil {
		return response.([]uint64), err
	}
	return []uint64{}, err
}
