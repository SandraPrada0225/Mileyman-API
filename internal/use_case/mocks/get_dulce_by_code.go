package mocks

import (
	"Mileyman-API/internal/domain/dto/query"

	"github.com/stretchr/testify/mock"
)

type MockGetDulceByCode struct {
	mock.Mock
}

func (mock *MockGetDulceByCode) Execute(codigo string) (query.DetalleDulce, error) {
	response := mock.Called(codigo)
	dulce := response.Get(0)
	err := response.Error(1)
	if err != nil {
		return query.DetalleDulce{}, err
	}
	return dulce.(query.DetalleDulce), nil
}
