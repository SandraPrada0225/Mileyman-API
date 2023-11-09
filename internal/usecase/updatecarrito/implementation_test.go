package updatecarrito

import (
	"errors"
	"reflect"
	"testing"

	"Mileyman-API/internal/domain/dto/contracts/updatecarrito"
	"Mileyman-API/internal/domain/dto/responses"
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
	movements := getMockMovements()
	carritoDulce := getMockCarritoDulce()
	dulce1 := getMockDulce1()
	dulce2 := getMockDulce2()

	mockCarritosProvider.On("GetByID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(carritoDulce, true, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[0].DulceID).Return(dulce1, nil)
	mockCarritosProvider.On("AddDulceInCarrito", getMockCarritoDulceUpdated()).Return(nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[1].DulceID).Return(entities.CarritoDulce{}, false, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[1].DulceID).Return(dulce2, nil)
	mockCarritosProvider.On("AddDulceInCarrito", getMockCarritoDulce2()).Return(nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[2].DulceID).Return(carritoDulce, true, nil)
	mockCarritosProvider.On("DeleteDulceInCarrito", carritoDulce).Return(nil)

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 3)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetByID", 2)
	mockCarritosProvider.AssertNumberOfCalls(t, "AddDulceInCarrito", 2)
	mockCarritosProvider.AssertNumberOfCalls(t, "DeleteDulceInCarrito", 1)
}

func TestWhentDeleteFailedThenShouldError(t *testing.T) {
	initialize()
	movements := getMockMovements2()
	carritoDulce := getMockCarritoDulce2()

	mockCarritosProvider.On("GetByID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(entities.CarritoDulce{}, false, nil)
	mockCarritosProvider.On("DeleteDulceInCarrito", entities.CarritoDulce{}).Return(errors.New("No se encontró un detalle carrito_dulce con ese codigo. resource: carrito"))

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse2(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "DeleteDulceInCarrito", 1)
}

func TestWhentAddDulceFailedThenShouldUnitLimitExceded(t *testing.T) {
	initialize()
	movements := getMockMovements3()
	carritoDulce := getMockCarritoDulce()
	dulce1 := getMockDulce1()

	mockCarritosProvider.On("GetByID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(carritoDulce, true, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[0].DulceID).Return(dulce1, nil)

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse3(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetByID", 1)
}

func TestWhenGetByIDFailedThenShouldNotFoundError(t *testing.T) {
	initialize()
	movements := getMockMovements4()
	carritoDulce := getMockCarritoDulce()

	mockCarritosProvider.On("GetByID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(entities.CarritoDulce{}, false, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[0].DulceID).Return(entities.Dulce{}, errors.New("No se encontró un dulce con ese codigo. resource: dulce"))

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse4(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 1)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetByID", 1)
}

func TestWhenGetDulceByCarritoIDAndDulceIDFailedThenShouldInternalServerError(t *testing.T) {
	initialize()
	movements := getMockMovements4()
	carritoDulce := getMockCarritoDulce()

	mockCarritosProvider.On("GetByID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(entities.CarritoDulce{}, false, errors.New("Ha ocurrido un error inesperado"))

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse5(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 1)
}

func TestWhentOneMovementErrorThenShouldOK(t *testing.T) {
	initialize()
	movements := getMockMovements5()
	carritoDulce := getMockCarritoDulce()
	dulce1 := getMockDulce1()
	dulce2 := getMockDulce2()

	mockCarritosProvider.On("GetByID", carritoDulce.CarritoID).Return(getMockCarrito(), nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[0].DulceID).Return(carritoDulce, true, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[0].DulceID).Return(dulce1, nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[1].DulceID).Return(entities.CarritoDulce{}, false, nil)
	mockDulcesProvider.On("GetByID", movements.Movements[1].DulceID).Return(dulce2, nil)
	mockCarritosProvider.On("AddDulceInCarrito", getMockCarritoDulce2()).Return(nil)

	mockCarritosProvider.On("GetDulceByCarritoIDAndDulceID", carritoDulce.CarritoID, movements.Movements[2].DulceID).Return(carritoDulce, true, nil)
	mockCarritosProvider.On("DeleteDulceInCarrito", carritoDulce).Return(nil)

	queryResponse, err := useCase.Execute(carritoDulce.CarritoID, movements)

	assert.NoError(t, err)
	assert.Equal(t, getMockExpectedResponse6(), queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetDulceByCarritoIDAndDulceID", 3)
	mockDulcesProvider.AssertNumberOfCalls(t, "GetByID", 2)
	mockCarritosProvider.AssertNumberOfCalls(t, "AddDulceInCarrito", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "DeleteDulceInCarrito", 1)
}

func TestWhenGetCarritoByCarritoIDFailedThenShouldNotFoundError(t *testing.T) {
	initialize()
	movements := getMockMovements4()

	mockCarritosProvider.On("GetByID", mockCarritoID2).Return(entities.Carrito{}, database.NewNotFoundError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID2, movements)

	assert.Error(t, err)

	typeErr := reflect.TypeOf(err).String()

	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
}

func TestWhenGetCarritoByCarritoIDFailedThenShouldInternalServerError(t *testing.T) {
	initialize()
	movements := getMockMovements4()

	mockCarritosProvider.On("GetByID", mockCarritoID2).Return(entities.Carrito{}, database.NewInternalServerError("error"))

	queryResponse, err := useCase.Execute(mockCarritoID2, movements)

	assert.Error(t, err)

	typeErr := reflect.TypeOf(err).String()

	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, queryResponse)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
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

func getMockCarritoDulce() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		ID:        1,
		CarritoID: 1,
		DulceID:   1,
		Unidades:  2,
		Subtotal:  5900,
	}
	return
}

func getMockCarritoDulce2() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		CarritoID: 1,
		DulceID:   2,
		Unidades:  2,
		Subtotal:  2000,
	}
	return
}

func getMockCarritoDulceUpdated() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		ID:        1,
		CarritoID: 1,
		DulceID:   1,
		Unidades:  4,
		Subtotal:  11800,
	}
	return
}

func getMockMovements() (movements updatecarrito.Body) {
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

func getMockMovements2() (movements updatecarrito.Body) {
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

func getMockMovements3() (movements updatecarrito.Body) {
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

func getMockMovements4() (movements updatecarrito.Body) {
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

func getMockMovements5() (movements updatecarrito.Body) {
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

func getMockExpectedResponse() responses.MovementsResult {
	return responses.MovementsResult{
		Result: []responses.MovementResult{
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

func getMockExpectedResponse2() responses.MovementsResult {
	return responses.MovementsResult{
		Result: []responses.MovementResult{
			{
				Movement: 0,
				DulceID:  1,
				Result:   "Error",
				Error:    "No se encontró un detalle carrito_dulce con ese codigo. resource: carrito",
			},
		},
	}
}

func getMockExpectedResponse3() responses.MovementsResult {
	return responses.MovementsResult{
		Result: []responses.MovementResult{
			{
				Movement: 0,
				DulceID:  1,
				Result:   "Error",
				Error:    "las unidades requeridad exceden las disponibles",
			},
		},
	}
}

func getMockExpectedResponse4() responses.MovementsResult {
	return responses.MovementsResult{
		Result: []responses.MovementResult{
			{
				Movement: 0,
				DulceID:  3,
				Result:   "Error",
				Error:    "No se encontró un dulce con ese codigo. resource: dulce",
			},
		},
	}
}

func getMockExpectedResponse5() responses.MovementsResult {
	return responses.MovementsResult{
		Result: []responses.MovementResult{
			{
				Movement: 0,
				DulceID:  3,
				Result:   "Error",
				Error:    "Ha ocurrido un error inesperado",
			},
		},
	}
}

func getMockExpectedResponse6() responses.MovementsResult {
	return responses.MovementsResult{
		Result: []responses.MovementResult{
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
		Subtotal:    0,
		PrecioTotal: 0,
		Descuento:   0,
		Envio:       0,
	}
}
