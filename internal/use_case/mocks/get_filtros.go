package mocks

import (
	"Mileyman-API/internal/domain/dto/query"

	"github.com/stretchr/testify/mock"
)

type MockGetFiltros struct {
	mock.Mock
}

func (mock *MockGetFiltros) Execute() (query.GetFiltros, error) {
	response := mock.Called()
	filtros := response.Get(0)
	err := response.Error(1)
	if err != nil {
		return query.GetFiltros{}, err
	}
	return filtros.(query.GetFiltros), nil
}
