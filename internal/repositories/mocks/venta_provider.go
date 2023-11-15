package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockVentaProvider struct {
	mock.Mock
}

func (m *MockVentaProvider) Create(venta *entities.Venta) error {
	args := m.Called(venta)
	venta.ID = args.Get(0).(uint64)
	err := args.Error(1)
	return err
}
