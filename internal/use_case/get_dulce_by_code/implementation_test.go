package getdulcebycode

import (
	"testing"

	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

var (
	useCase           Implementation
	mockDulceProvider *mocks.MockDulceProvider
)

func initialize() {
	mockDulceProvider = new(mocks.MockDulceProvider)
	useCase = Implementation{
		DulcesProvider: mockDulceProvider,
	}
}

func TestWhenSuccesfullReturnDulce(t *testing.T) {
	initialize()
	expectedDulce := GetDulce()
	mockDulceProvider.On("GetByCode", expectedDulce.Codigo).Return(expectedDulce, nil)
	dulce, err := useCase.Execute(expectedDulce.Codigo)

	assert.NoError(t, err)
	assert.Equal(t, expectedDulce, dulce)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
}

func TestWhenDulceNotFoundReturnNotFoundError(t *testing.T) {
	initialize()
	expectedDulce := GetDulce()
	mockDulceProvider.On("GetByCode", expectedDulce.Codigo).Return(expectedDulce, nil)
	dulce, err := useCase.Execute(expectedDulce.Codigo)

	assert.NoError(t, err)
	assert.Equal(t, expectedDulce, dulce)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
}

func GetDulce() (dulce entities.Dulce) {
	dulce = entities.Dulce{
		ID:             2,
		Nombre:         "Chocolatina",
		PresentacionID: 1,
		Descripcion:    "Deliciosa chocolatina que se derrite en tu boca",
		Imagen:         "imagen",
		Disponibles:    100,
		Precio:         1000,
		Peso:           40,
		MarcaID:        1,
		Codigo:         "2",
	}
	return
}
