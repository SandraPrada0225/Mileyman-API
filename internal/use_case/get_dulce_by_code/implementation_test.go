package getdulcebycode

import (
	"reflect"
	"testing"

	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

var (
	useCase               Implementation
	mockDulceProvider     *mocks.MockDulceProvider
	mockCategoriaProvider *mocks.MockCategoriasProvider
)

func initialize() {
	mockDulceProvider = new(mocks.MockDulceProvider)
	mockCategoriaProvider = new(mocks.MockCategoriasProvider)
	useCase = Implementation{
		DulcesProvider:     mockDulceProvider,
		CategoriasProvider: mockCategoriaProvider,
	}
}

func TestWhenSuccesfullReturnDulce(t *testing.T) {
	initialize()
	partialResponse := GetPartialResponse()
	mockCategorias := GetMockCategorias()
	expectedResponse := partialResponse
	expectedResponse.Categorias = mockCategorias

	mockDulceProvider.On("GetByCode", partialResponse.Codigo).Return(partialResponse, nil)
	mockCategoriaProvider.On("GetCategoriasByDulceID", partialResponse.ID).Return(mockCategorias, nil)
	dulce, err := useCase.Execute(partialResponse.Codigo)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, dulce)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 1)
}

func TestWhenGetDulceWentWrongShouldReturnInternalServerError(t *testing.T) {
	initialize()
	expectedDulce := GetPartialResponse()
	mockDulceProvider.On("GetByCode", expectedDulce.Codigo).Return(query.DetalleDulce{}, database.NewInternalServerError("error"))
	dulce, err := useCase.Execute(expectedDulce.Codigo)

	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, dulce)
	assert.Equal(t, "database.InternalServerError", errType)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 0)
}

func TestWhenDulceNotFoundThenShouldReturnNotFoundError(t *testing.T) {
	initialize()
	expectedDulce := GetPartialResponse()
	mockDulceProvider.On("GetByCode", expectedDulce.Codigo).Return(query.DetalleDulce{}, database.NewNotFoundError("error"))
	dulce, err := useCase.Execute(expectedDulce.Codigo)

	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, dulce)
	assert.Equal(t, "database.NotFoundError", errType)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 0)
}

func TestWhenGetCategoriasWentWrongShouldReturnInternalServer(t *testing.T) {
	initialize()
	partialResponse := GetPartialResponse()
	mockCategorias := GetMockCategorias()
	expectedResponse := partialResponse
	expectedResponse.Categorias = mockCategorias

	mockDulceProvider.On("GetByCode", partialResponse.Codigo).Return(partialResponse, nil)
	mockCategoriaProvider.On("GetCategoriasByDulceID", partialResponse.ID).Return([]entities.Categoria{}, database.NewInternalServerError("error"))
	response, err := useCase.Execute(partialResponse.Codigo)

	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, response)
	assert.Equal(t, "database.InternalServerError", errType)
	mockDulceProvider.AssertNumberOfCalls(t, "GetByCode", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 1)
}

func GetPartialResponse() query.DetalleDulce {
	return query.DetalleDulce{
		ID:           2,
		Nombre:       "Chocolatina",
		Descripcion:  "Deliciosa chocolatina que se derrite en tu boca",
		Imagen:       "imagen",
		Disponibles:  100,
		PrecioUnidad: 1000,
		Peso:         40,
		Codigo:       "2",
		Presentacion: entities.Presentacion{
			ID:     1,
			Nombre: "Empaque",
		},
		Marca: entities.Marca{
			ID:     2,
			Nombre: "Jet",
		},
	}
}

func GetMockCategorias() (categorias []entities.Categoria) {
	categorias = []entities.Categoria{
		{
			ID:     1,
			Nombre: "Gomitas",
		},
		{
			ID:     2,
			Nombre: "Chocolates",
		},
	}
	return
}
