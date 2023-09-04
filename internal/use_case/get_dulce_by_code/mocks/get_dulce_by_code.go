package mocks

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockGetDulceByCode struct {
	mock.Mock
}

func (mock *MockGetDulceByCode) Execute(codigo string) (entities.Dulce, error) {
		response := mock.Called(codigo)
		dulce := response.Get(0)
		err := response.Error(1)
		if(err != nil){
			return entities.Dulce{}, err
		}
		return dulce.(entities.Dulce), nil
}
