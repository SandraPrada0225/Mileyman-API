package mocks

import (
	"Mileyman-API/internal/domain/dto/query"

	"github.com/stretchr/testify/mock"
)

type MockGetCarritoByID struct {
	mock.Mock
}

func (mock *MockGetCarritoByID) Execute(id uint64) (query.GetDetalleCarrito, error) {
	response := mock.Called(id)
	dulce := response.Get(0)
	err := response.Error(1)
	if err != nil {
		return query.GetDetalleCarrito{}, err
	}
	return dulce.(query.GetDetalleCarrito), nil
}
