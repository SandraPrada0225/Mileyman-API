package mocks

import (
	"Mileyman-API/internal/domain/dto/responses"

	"github.com/stretchr/testify/mock"
)

type MockGetDulceByCode struct {
	mock.Mock
}

func (mock *MockGetDulceByCode) Execute(codigo string) (responses.DetalleDulce, error) {
	response := mock.Called(codigo)
	dulce := response.Get(0)
	err := response.Error(1)
	if err != nil {
		return responses.DetalleDulce{}, err
	}
	return dulce.(responses.DetalleDulce), nil
}
