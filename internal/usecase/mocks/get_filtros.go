package mocks

import (
	"Mileyman-API/internal/domain/dto/responses"

	"github.com/stretchr/testify/mock"
)

type MockGetFiltros struct {
	mock.Mock
}

func (mock *MockGetFiltros) Execute() (responses.GetFiltros, error) {
	response := mock.Called()
	filtros := response.Get(0)
	err := response.Error(1)
	if err != nil {
		return responses.GetFiltros{}, err
	}
	return filtros.(responses.GetFiltros), nil
}
