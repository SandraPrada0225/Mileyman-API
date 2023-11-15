package mocks

import (
	"Mileyman-API/internal/domain/dto/responses"

	"github.com/stretchr/testify/mock"
)

type MockGetCarritoByID struct {
	mock.Mock
}

func (mock *MockGetCarritoByID) Execute(id uint64) (responses.GetDetalleCarrito, error) {
	response := mock.Called(id)
	dulce := response.Get(0)
	err := response.Error(1)
	if err != nil {
		return responses.GetDetalleCarrito{}, err
	}
	return dulce.(responses.GetDetalleCarrito), nil
}
