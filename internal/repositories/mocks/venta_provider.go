package mocks

import (
	"Mileyman-API/internal/domain/dto/responses"
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

func (m *MockVentaProvider) GetListByUserID(userID uint64) (responses.GetPurchaseList, error) {
	args := m.Called(userID)
	response := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return responses.GetPurchaseList{}, err
	}
	return response.(responses.GetPurchaseList), err
}
