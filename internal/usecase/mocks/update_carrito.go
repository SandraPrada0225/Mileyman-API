package mocks

import (
	"Mileyman-API/internal/domain/dto/command/updatecarrito"
	"Mileyman-API/internal/domain/dto/responses"

	"github.com/stretchr/testify/mock"
)

type MockUpdateCarrito struct {
	mock.Mock
}

func (mock *MockUpdateCarrito) Execute(carritoID uint64, movements updatecarrito.Body) (responses.MovementsResult, error) {
	response := mock.Called(carritoID, movements)
	movementsResult := response.Get(0)
	err := response.Error(1)
	if err != nil {
		return responses.MovementsResult{}, err
	}
	return movementsResult.(responses.MovementsResult), err
}
