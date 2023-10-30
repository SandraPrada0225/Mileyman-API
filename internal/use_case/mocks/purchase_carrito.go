package mocks

import (
	"Mileyman-API/internal/domain/dto/command"

	"github.com/stretchr/testify/mock"
)

type MockPurchaseCarrito struct {
	mock.Mock
}

func (m *MockPurchaseCarrito) Execute(command command.PurchaseCarritoCommand) (uint64, error) {
	args := m.Called(command)
	carritoID := args.Get(0).(uint64)
	err := args.Error(1)
	return carritoID, err
}
