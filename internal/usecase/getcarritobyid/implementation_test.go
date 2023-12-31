package getcarritobyid

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
	mockCarritosProvider   *mocks.MockCarritoProvider
	mockCategoriasProvider *mocks.MockCategoriaProvider
	mockDulcesProvider     *mocks.MockDulceProvider
)

const (
	mockCarritoID = uint64(2321)
	mockDulceID1  = uint64(1231)
	mockDulceID2  = uint64(2321)
)

func TestWhenEverythingWentSucessfullyThenShouldReturnCarrito(t *testing.T) {
	useCase := getImplementation()

	mockCarritosProvider.On("GetByID", mockCarritoID).
		Return(getMockCarrito(), nil)
	mockDulcesProvider.On("GetDulcesListByCarritoID", mockCarritoID).
		Return(getMockDulcesInCarritoList(), nil)
	mockDulcesProvider.On("GetDetailByID", mockDulceID1).
		Return(getPartialResponse1(), nil)
	mockDulcesProvider.On("GetDetailByID", mockDulceID2).
		Return(getPartialResponse2(), nil)
	mockCategoriasProvider.On("GetCategoriasByDulceID", mockDulceID1).
		Return(getMockCategorias1(), nil)
	mockCategoriasProvider.On("GetCategoriasByDulceID", mockDulceID2).
		Return(getMockCategorias2(), nil)

	queryResponse, err := useCase.Execute(mockCarritoID)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDulcesListByCarritoID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDetailByID", 2)
	mockCategoriasProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 2)
}

func TestWhenCarritoWasNotFoundThenShouldReturnNotFoundError(t *testing.T) {
	useCase := getImplementation()

	mockCarritosProvider.On("GetByID", mockCarritoID).
		Return(entities.Carrito{}, database.NewNotFoundError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID)

	assert.Error(t, err)

	typeErr := reflect.TypeOf(err).String()

	assert.Empty(t, queryResponse)
	assert.Equal(t, "database.NotFoundError", typeErr)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDulcesListByCarritoID", 0)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDetailByID", 0)
	mockCategoriasProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 0)
}

func TestWhenCarritoWentWrongThenShouldReturnInternalServerError(t *testing.T) {
	useCase := getImplementation()

	mockCarritosProvider.On("GetByID", mockCarritoID).
		Return(entities.Carrito{}, database.NewInternalServerError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID)

	assert.Error(t, err)

	typeErr := reflect.TypeOf(err).String()

	assert.Empty(t, queryResponse)
	assert.Equal(t, "database.InternalServerError", typeErr)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDulcesListByCarritoID", 0)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDetailByID", 0)
	mockCategoriasProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 0)
}

func TestWhenGetDulcesCarritoWentWrongThenShouldReturnCarrito(t *testing.T) {
	useCase := getImplementation()

	mockCarritosProvider.On("GetByID", mockCarritoID).
		Return(getMockCarrito(), nil)
	mockDulcesProvider.On("GetDulcesListByCarritoID", mockCarritoID).
		Return(getMockDulcesInCarritoList(), database.NewInternalServerError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Empty(t, queryResponse)
	assert.Equal(t, "database.InternalServerError", typeErr)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDulcesListByCarritoID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDetailByID", 0)
	mockCategoriasProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 0)
}

func TestWhenGetDulceByIDWentWrongThenShouldReturnInternalServerError(t *testing.T) {
	useCase := getImplementation()

	mockCarritosProvider.On("GetByID", mockCarritoID).
		Return(getMockCarrito(), nil)
	mockDulcesProvider.On("GetDulcesListByCarritoID", mockCarritoID).
		Return(getMockDulcesInCarritoList(), nil)
	mockDulcesProvider.On("GetDetailByID", mockDulceID1).
		Return(responses.DetalleDulce{}, database.NewInternalServerError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Empty(t, queryResponse)
	assert.Equal(t, "database.InternalServerError", typeErr)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDulcesListByCarritoID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDetailByID", 1)
	mockCategoriasProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 0)
}

func TestWhenGeCategoriasByDulceIDWentWrongThenShouldReturnInternalServerError(t *testing.T) {
	useCase := getImplementation()

	mockCarritosProvider.On("GetByID", mockCarritoID).
		Return(getMockCarrito(), nil)
	mockDulcesProvider.On("GetDulcesListByCarritoID", mockCarritoID).
		Return(getMockDulcesInCarritoList(), nil)
	mockDulcesProvider.On("GetDetailByID", mockDulceID1).
		Return(getPartialResponse1(), nil)
	mockCategoriasProvider.On("GetCategoriasByDulceID", mockDulceID1).
		Return([]entities.Categoria{}, database.NewInternalServerError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Empty(t, queryResponse)
	assert.Equal(t, "database.InternalServerError", typeErr)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDulcesListByCarritoID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetDetailByID", 1)
	mockCategoriasProvider.AssertNumberOfCalls(t, "GetCategoriasByDulceID", 1)
}

func getImplementation() Implementation {
	mockCarritosProvider = new(mocks.MockCarritoProvider)
	mockCategoriasProvider = new(mocks.MockCategoriaProvider)
	mockDulcesProvider = new(mocks.MockDulceProvider)
	return Implementation{
		CarritoProvider:    mockCarritosProvider,
		CategoriasProvider: mockCategoriasProvider,
		DulcesProvider:     mockDulcesProvider,
	}
}

func getMockCarrito() entities.Carrito {
	return entities.Carrito{
		ID:          mockCarritoID,
		Subtotal:    5,
		PrecioTotal: 100,
		Descuento:   5,
		Envio:       5,
	}
}

func getMockDulcesInCarritoList() []entities.CarritoDulce {
	return []entities.CarritoDulce{
		{
			ID:        1,
			DulceID:   mockDulceID1,
			CarritoID: mockCarritoID,
			Unidades:  2,
			Subtotal:  2000,
		},
		{
			ID:        2,
			DulceID:   mockDulceID2,
			CarritoID: mockCarritoID,
			Unidades:  1,
			Subtotal:  1000,
		},
	}
}

func getPartialResponse1() responses.DetalleDulce {
	return responses.DetalleDulce{
		ID:           mockDulceID1,
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

func getPartialResponse2() responses.DetalleDulce {
	return responses.DetalleDulce{
		ID:           mockDulceID2,
		Nombre:       "Gomitas",
		Descripcion:  "Ositos de gomita con sabores explosivos",
		Imagen:       "imagen",
		Disponibles:  100,
		PrecioUnidad: 1000,
		Peso:         40,
		Codigo:       "2",
		Presentacion: entities.Presentacion{
			ID:     1,
			Nombre: "Paquete",
		},
		Marca: entities.Marca{
			ID:     2,
			Nombre: "Trululú",
		},
	}
}

func getMockCategorias1() (categorias []entities.Categoria) {
	categorias = []entities.Categoria{
		{
			ID:     2,
			Nombre: "Chocolates",
		},
	}
	return
}

func getMockCategorias2() (categorias []entities.Categoria) {
	categorias = []entities.Categoria{
		{
			ID:     1,
			Nombre: "Gomitas",
		},
	}
	return
}

func getMockExpectedResponse() responses.GetDetalleCarrito {
	return responses.GetDetalleCarrito{
		ID:          mockCarritoID,
		Subtotal:    5,
		PrecioTotal: 100,
		Descuento:   5,
		Envio:       5,
		DulcesList: []responses.DulceInCarrito{
			{
				DetalleDulce: responses.DetalleDulce{
					ID:           mockDulceID1,
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
					Categorias: []entities.Categoria{
						{
							ID:     2,
							Nombre: "Chocolates",
						},
					},
				},
				Unidades: 2,
				Subtotal: 2000,
			},
			{
				DetalleDulce: responses.DetalleDulce{
					ID:           mockDulceID2,
					Nombre:       "Gomitas",
					Descripcion:  "Ositos de gomita con sabores explosivos",
					Imagen:       "imagen",
					Disponibles:  100,
					PrecioUnidad: 1000,
					Peso:         40,
					Codigo:       "2",
					Presentacion: entities.Presentacion{
						ID:     1,
						Nombre: "Paquete",
					},
					Marca: entities.Marca{
						ID:     2,
						Nombre: "Trululú",
					},
					Categorias: []entities.Categoria{
						{
							ID:     1,
							Nombre: "Gomitas",
						},
					},
				},
				Unidades: 1,
				Subtotal: 1000,
			},
		},
	}
}
