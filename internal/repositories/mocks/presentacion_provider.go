package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockPresentacionProvider struct {
	mock.Mock 
}

func (mock *MockPresentacionProvider) GetAll() ([]entities.Presentacion, error) {
	args := mock.Called()
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.([]entities.Presentacion), err
	}
	return []entities.Presentacion{}, err
}
