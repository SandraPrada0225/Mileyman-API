package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockCarritoProvider struct {
	mock.Mock
}

func (mock *MockCarritoProvider) GetCarritoByCarritoID(carritoID uint64) (entities.Carrito, error) {
	args := mock.Called(carritoID)
	response := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return entities.Carrito{}, err
	}
	return response.(entities.Carrito), nil
}

func (mock *MockCarritoProvider) GetDulceByCarritoIDAndDulceID(carritoID uint64, dulceID uint64) (carritoDulce entities.CarritoDulce, exists bool, err error) {
	args := mock.Called(carritoID, dulceID)
	response1 := args.Get(0)
	response2 := args.Bool(1)
	err = args.Error(2)
	if err != nil {
		return entities.CarritoDulce{}, false, err
	}
	return response1.(entities.CarritoDulce), response2, nil
}

func (mock *MockCarritoProvider) AddDulceInCarrito(carritoDulce entities.CarritoDulce) (err error) {
	args := mock.Called(carritoDulce)
	err = args.Error(0)
	if err != nil {
		return err
	}
	return nil
}

func (mock *MockCarritoProvider) DeleteDulceInCarrito(carritoDulce entities.CarritoDulce) (err error) {
	args := mock.Called(carritoDulce)
	err = args.Error(0)
	if err != nil {
		return err
	}
	return nil
}
