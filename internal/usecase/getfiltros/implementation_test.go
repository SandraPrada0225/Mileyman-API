package getfiltros

import (
	"reflect"
	"testing"

	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

var (
	useCase                  Implementation
	mockCategoriaProvider    *mocks.MockCategoriaProvider
	mockMarcaProvider        *mocks.MockMarcaProvider
	mockPresentacionProvider *mocks.MockPresentacionProvider
)

func initialize() {
	mockCategoriaProvider = new(mocks.MockCategoriaProvider)
	mockMarcaProvider = new(mocks.MockMarcaProvider)
	mockPresentacionProvider = new(mocks.MockPresentacionProvider)

	useCase = Implementation{
		CategoriasProvider:     mockCategoriaProvider,
		MarcasProvider:         mockMarcaProvider,
		PresentacionesProvider: mockPresentacionProvider,
	}
}

func TestWhenSuccesfullReturnAll(t *testing.T) {
	initialize()
	expectedFiltros := GetFiltros()

	mockCategoriaProvider.On("GetAll").Return(expectedFiltros.Categorias, nil)
	mockMarcaProvider.On("GetAll").Return(expectedFiltros.Marcas, nil)
	mockPresentacionProvider.On("GetAll").Return(expectedFiltros.Presentaciones, nil)

	filtros, err := useCase.Execute()

	assert.NoError(t, err)
	assert.Equal(t, expectedFiltros, filtros)
	mockMarcaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockPresentacionProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetAll", 1)
}

func TestWentWrongGetMarcasRun(t *testing.T) {
	initialize()

	mockMarcaProvider.On("GetAll").Return([]entities.Marca{}, database.NewInternalServerError(""))

	filtros, err := useCase.Execute()

	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, filtros)
	assert.Equal(t, "database.InternalServerError", errType)
	mockMarcaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockPresentacionProvider.AssertNumberOfCalls(t, "GetAll", 0)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetAll", 0)
}

func TestWentWrongGetPresentacionesRun(t *testing.T) {
	initialize()
	expectedFiltros := GetFiltros()
	mockMarcaProvider.On("GetAll").Return(expectedFiltros.Marcas, nil)
	mockPresentacionProvider.On("GetAll").Return([]entities.Presentacion{}, database.NewInternalServerError(""))

	filtros, err := useCase.Execute()

	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, filtros)
	assert.Equal(t, "database.InternalServerError", errType)
	mockMarcaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockPresentacionProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetAll", 0)
}

func TestWentWrongGetCategoriasRun(t *testing.T) {
	initialize()
	expectedFiltros := GetFiltros()
	mockMarcaProvider.On("GetAll").Return(expectedFiltros.Marcas, nil)
	mockPresentacionProvider.On("GetAll").Return(expectedFiltros.Presentaciones, nil)
	mockCategoriaProvider.On("GetAll").Return([]entities.Categoria{}, database.NewInternalServerError(""))

	filtros, err := useCase.Execute()

	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Empty(t, filtros)
	assert.Equal(t, "database.InternalServerError", errType)
	mockMarcaProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockPresentacionProvider.AssertNumberOfCalls(t, "GetAll", 1)
	mockCategoriaProvider.AssertNumberOfCalls(t, "GetAll", 1)
}

func GetFiltros() (filtros responses.GetFiltros) {
	filtros = responses.GetFiltros{
		Categorias: []entities.Categoria{
			{
				ID:     1,
				Nombre: "Gomitas",
			},
			{
				ID:     2,
				Nombre: "Chupetes",
			},
		},
		Marcas: []entities.Marca{
			{
				ID:     1,
				Nombre: "Trululu",
			},
			{
				ID:     2,
				Nombre: "Jet",
			},
		},
		Presentaciones: []entities.Presentacion{
			{
				ID:     1,
				Nombre: "Caja",
			},
			{
				ID:     2,
				Nombre: "Bolsa",
			},
		},
	}
	return
}
