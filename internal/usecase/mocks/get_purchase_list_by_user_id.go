package mocks

import (
	"Mileyman-API/internal/domain/dto/responses"

	"github.com/stretchr/testify/mock"
)

type MockGetPurchaseListByUSerID struct {
	mock.Mock
}

func (m *MockGetPurchaseListByUSerID) Execute(userID uint64) (responses.GetPurchaseList, error) {
	args := m.Called(userID)
	response := args.Get(0)
	err := args.Error(1)
	if err != nil {
		return responses.GetPurchaseList{}, err
	}
	return response.(responses.GetPurchaseList), err
}
