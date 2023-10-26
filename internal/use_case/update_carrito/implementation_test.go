package updatecarrito

import (
	"errors"
	"reflect"
	"testing"

	updatecarrito "Mileyman-API/internal/domain/dto/command/update_carrito"
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

const (
	mockCarritoID  = uint64(1)
	mockCarritoID2 = uint64(2)
)

var (
	useCase              Implementation
	mockCarritosProvider *mocks.MockCarritoProvider
	mockDulcesProvider   *mocks.MockDulceProvider
)

func initialize() {
	mockCarritosProvider = new(mocks.MockCarritoProvider)
	mockDulcesProvider = new(mocks.MockDulceProvider)

	useCase = Implementation{
		CarritosProvider: mockCarritosProvider,
		DulcesProvider:   mockDulcesProvider,
	}
}

func TestWhentSucessfullyThenShouldOK(t *testing.T) {
	initialize()
	movements := GetMockMovements()
	carritoDulce := GetMockCarritoDulce()
	dulce1 := getMockDulce1()
	dulce2 := getMockDulce2()

	mockCarritosProvider.On("GetCarritoByCarritoID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	// Updated
	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(carritoDulce, true, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[0].DulceID).Return(dulce1, nil)
	mockCarritosProvider.On("AddDulceInCarrito", GetMockCarritoDulceUpdated()).Return(nil)

	// created
	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[1].DulceID).Return(entities.CarritoDulce{}, false, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[1].DulceID).Return(dulce2, nil)
	mockCarritosProvider.On("AddDulceInCarrito", GetMockCarritoDulce2()).Return(nil)

	// Deleted
	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[2].DulceID).Return(carritoDulce, true, nil)
	mockCarritosProvider.On("DeleteDulceInCarrito", carritoDulce).Return(nil)

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetCarritoByCarritoID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 3)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetByID", 2)
	mockCarritosProvider.AssertNumberOfCalls(t, "AddDulceInCarrito", 2)
	mockCarritosProvider.AssertNumberOfCalls(t, "DeleteDulceInCarrito", 1)
}

func TestWhentDeleteFailedThenShouldError(t *testing.T) {
	initialize()
	movements := GetMockMovements2()
	carritoDulce := GetMockCarritoDulce2()

	mockCarritosProvider.On("GetCarritoByCarritoID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)
	// Deleted
	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(entities.CarritoDulce{}, false, nil)
	mockCarritosProvider.On("DeleteDulceInCarrito", entities.CarritoDulce{}).Return(errors.New("No se encontró un detalle carrito_dulce con ese codigo. resource: carrito"))

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse2(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetCarritoByCarritoID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "DeleteDulceInCarrito", 1)
}

func TestWhentAddDulceFailedThenShouldUnitLimitExceded(t *testing.T) {
	initialize()
	movements := GetMockMovements3()
	carritoDulce := GetMockCarritoDulce()
	dulce1 := getMockDulce1()

	mockCarritosProvider.On("GetCarritoByCarritoID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	// Updated
	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(carritoDulce, true, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[0].DulceID).Return(dulce1, nil)

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse3(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetCarritoByCarritoID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetByID", 1)
}

func TestWhenGetByIDFailedThenShouldNotFoundError(t *testing.T) {
	initialize()
	movements := GetMockMovements4()
	carritoDulce := GetMockCarritoDulce()

	mockCarritosProvider.On("GetCarritoByCarritoID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(entities.CarritoDulce{}, false, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[0].DulceID).Return(entities.Dulce{}, errors.New("No se encontró un dulce con ese codigo. resource: dulce"))

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse4(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetCarritoByCarritoID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetByID", 1)
}

func TestWhenGetDulceByCarritoIDAndDulceIDFailedThenShouldInternalServerError(t *testing.T) {
	initialize()
	movements := GetMockMovements4()
	carritoDulce := GetMockCarritoDulce()

	mockCarritosProvider.On("GetCarritoByCarritoID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(entities.CarritoDulce{}, false, errors.New("Ha ocurrido un error inesperado"))

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse5(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetCarritoByCarritoID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 1)
}

func TestWhentOneMovementErrorThenShouldOK(t *testing.T) {
	initialize()
	movements := GetMockMovements5()
	carritoDulce := GetMockCarritoDulce()
	dulce1 := getMockDulce1()
	dulce2 := getMockDulce2()

	mockCarritosProvider.On("GetCarritoByCarritoID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	// Updated
	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(carritoDulce, true, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[0].DulceID).Return(dulce1, nil)
	// created
	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[1].DulceID).Return(entities.CarritoDulce{}, false, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[1].DulceID).Return(dulce2, nil)
	mockCarritosProvider.On("AddDulceInCarrito", GetMockCarritoDulce2()).Return(nil)

	// Deleted
	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[2].DulceID).Return(carritoDulce, true, nil)
	mockCarritosProvider.On("DeleteDulceInCarrito", carritoDulce).Return(nil)

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse6(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetCarritoByCarritoID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 3)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetByID", 2)
	mockCarritosProvider.AssertNumberOfCalls(t, "AddDulceInCarrito", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "DeleteDulceInCarrito", 1)
}

func TestWhenGetCarritoByCarritoIDFailedThenShouldNotFoundError(t *testing.T) {
	initialize()
	movements := GetMockMovements4()

	mockCarritosProvider.On("GetCarritoByCarritoID", mockCarritoID2).Return(entities.Carrito{}, database.NewNotFoundError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID2, movements)

	assert.Error(t, err)

	typeErr := reflect.TypeOf(err).String()

	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetCarritoByCarritoID", 1)
}

func TestWhenGetCarritoByCarritoIDFailedThenShouldInternalServerError(t *testing.T) {
	initialize()
	movements := GetMockMovements4()

	mockCarritosProvider.On("GetCarritoByCarritoID", mockCarritoID2).Return(entities.Carrito{}, database.NewInternalServerError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID2, movements)

	assert.Error(t, err)

	typeErr := reflect.TypeOf(err).String()

	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetCarritoByCarritoID", 1)
}

func getMockDulce1() (Dulce entities.Dulce) {
	Dulce = entities.Dulce{
		ID:           1,
		Nombre:       "Gomas Clasicas",
		Descripcion:  "Gomas clasicas con sabores surtidos",
		Imagen:       "imagen",
		Disponibles:  100,
		PrecioUnidad: 2950,
		Peso:         80,
		Codigo:       "1A",
	}
	return
}

func getMockDulce2() (Dulce entities.Dulce) {
	Dulce = entities.Dulce{
		ID:           2,
		Nombre:       "Chocolatina",
		Descripcion:  "Deliciosa chocolatina que se derrite en tu boca",
		Imagen:       "imagen",
		Disponibles:  100,
		PrecioUnidad: 1000,
		Peso:         40,
		Codigo:       "1B",
	}
	return
}

func GetMockCarritoDulce() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		ID:        1,
		CarritoID: 1,
		DulceID:   1,
		Unidades:  2,
		Subtotal:  5900,
	}
	return
}

func GetMockCarritoDulce2() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		CarritoID: 1,
		DulceID:   2,
		Unidades:  2,
		Subtotal:  2000,
	}
	return
}

func GetMockCarritoDulceUpdated() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		ID:        1,
		CarritoID: 1,
		DulceID:   1,
		Unidades:  4,
		Subtotal:  11800,
	}
	return
}

func GetMockMovements() (movements updatecarrito.Body) {
	movements = updatecarrito.Body{
		Movements: []updatecarrito.Movement{
			{
				DulceID:  1,
				Unidades: 4,
			},
			{
				DulceID:  2,
				Unidades: 2,
			},
			{
				DulceID:  1,
				Unidades: 0,
			},
		},
	}
	return
}

func GetMockMovements2() (movements updatecarrito.Body) {
	movements = updatecarrito.Body{
		Movements: []updatecarrito.Movement{
			{
				DulceID:  1,
				Unidades: 0,
			},
		},
	}
	return
}

func GetMockMovements3() (movements updatecarrito.Body) {
	movements = updatecarrito.Body{
		Movements: []updatecarrito.Movement{
			{
				DulceID:  1,
				Unidades: 200,
			},
		},
	}
	return
}

func GetMockMovements4() (movements updatecarrito.Body) {
	movements = updatecarrito.Body{
		Movements: []updatecarrito.Movement{
			{
				DulceID:  3,
				Unidades: 100,
			},
		},
	}
	return
}

func GetMockMovements5() (movements updatecarrito.Body) {
	movements = updatecarrito.Body{
		Movements: []updatecarrito.Movement{
			{
				DulceID:  1,
				Unidades: 200,
			},
			{
				DulceID:  2,
				Unidades: 2,
			},
			{
				DulceID:  1,
				Unidades: 0,
			},
		},
	}
	return
}

func getMockExpectedResponse() query.MovementsResult {
	return query.MovementsResult{
		Result: []query.MovementResult{
			{
				Movement: 0,
				DulceID:  1,
				Result:   "Updated",
				Error:    "",
			},
			{
				Movement: 1,
				DulceID:  2,
				Result:   "Created",
				Error:    "",
			},
			{
				Movement: 2,
				DulceID:  1,
				Result:   "Deleted",
				Error:    "",
			},
		},
	}
}

func getMockExpectedResponse2() query.MovementsResult {
	return query.MovementsResult{
		Result: []query.MovementResult{
			{
				Movement: 0,
				DulceID:  1,
				Result:   "Error",
				Error:    "No se encontró un detalle carrito_dulce con ese codigo. resource: carrito",
			},
		},
	}
}

func getMockExpectedResponse3() query.MovementsResult {
	return query.MovementsResult{
		Result: []query.MovementResult{
			{
				Movement: 0,
				DulceID:  1,
				Result:   "Error",
				Error:    "las unidades requeridad exceden las disponibles",
			},
		},
	}
}

func getMockExpectedResponse4() query.MovementsResult {
	return query.MovementsResult{
		Result: []query.MovementResult{
			{
				Movement: 0,
				DulceID:  3,
				Result:   "Error",
				Error:    "No se encontró un dulce con ese codigo. resource: dulce",
			},
		},
	}
}

func getMockExpectedResponse5() query.MovementsResult {
	return query.MovementsResult{
		Result: []query.MovementResult{
			{
				Movement: 0,
				DulceID:  3,
				Result:   "Error",
				Error:    "Ha ocurrido un error inesperado",
			},
		},
	}
}

func getMockExpectedResponse6() query.MovementsResult {
	return query.MovementsResult{
		Result: []query.MovementResult{
			{
				Movement: 0,
				DulceID:  1,
				Result:   "Error",
				Error:    "las unidades requeridad exceden las disponibles",
			},
			{
				Movement: 1,
				DulceID:  2,
				Result:   "Created",
				Error:    "",
			},
			{
				Movement: 2,
				DulceID:  1,
				Result:   "Deleted",
				Error:    "",
			},
		},
	}
}

func getMockCarrito() entities.Carrito {
	return entities.Carrito{
		ID:          mockCarritoID,
		SubTotal:    0,
		PrecioTotal: 0,
		Descuento:   0,
		Envio:       0,
	}
}
