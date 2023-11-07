package mocks

import (
	updatecarrito "Mileyman-API/internal/domain/dto/command/update_carrito"
	"Mileyman-API/internal/domain/dto/query"

	"github.com/stretchr/testify/mock"
)

type MockUpdateCarrito struct {
	mock.Mock
}

func (mock *MockUpdateCarrito) Execute(carritoID uint64, movements updatecarrito.Body) (query.MovementsResult, error) {
	response := mock.Called(carritoID, movements)
	movementsResult := response.Get(0)
	err := response.Error(1)
	if err != nil {
		return query.MovementsResult{}, err
	}
	return movementsResult.(query.MovementsResult), err
}
