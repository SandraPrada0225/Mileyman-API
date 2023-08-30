package mocks

//dado unos parametros devuelve la respuesta deseada para comprobar que el caso de uso reaccione bien nate las diferentes situaciones 

import (
	"Mileyman-API/internal/domain/entities"

	"github.com/stretchr/testify/mock"
)

type MockDulceProvider struct {
	mock.Mock //implementacion boba de la interface
}

func (mock *MockDulceProvider) GetByCode(codigo string) (entities.Dulce, error) {
	args := mock.Called(codigo)
	response := args.Get(0)
	err := args.Error(1)

	if response != nil {
		return response.(entities.Dulce), err
	}
	return entities.Dulce{}, err
}
