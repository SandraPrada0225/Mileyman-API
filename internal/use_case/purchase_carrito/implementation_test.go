package purchasecarrito

import (
	"reflect"
	"testing"

	estadoscarrito "Mileyman-API/internal/domain/constants/estados_carrito"
	mediosdepago "Mileyman-API/internal/domain/constants/medios_de_pago"
	"Mileyman-API/internal/domain/dto/command"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"
	"Mileyman-API/internal/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

var (
	mockCarritosProvider *mocks.MockCarritoProvider
	mockUsuariosProvider *mocks.MockUsuarioProvider
	mockVentasProvider   *mocks.MockVentaProvider
)

const (
	mockCarritoID  uint64 = 5
	mockCarritoID2 uint64 = 6
)

func TestWhenEverythingIsSuccessfullShouldReturnNewCarritoID(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarrito := getMockCarrito()
	mockUpdatedCarrito := getMockUpdatedCarrito()
	mockVentaToCreate := getMockVentaToCreate()
	mockUsuario := getMockUsuario()
	mockUpdatedUsuario := getMockUpdatedUsuario()
	mockEmptyCarrito := getMockNewCarrito()

	mockCarritosProvider.On("GetByID", command.CarritoID).Return(mockCarrito, nil)
	mockUsuariosProvider.On("GetByID", command.CompradorID).Return(mockUsuario, nil)
	mockCarritosProvider.On("Save", &mockUpdatedCarrito).Return(uint64(0), nil)
	mockVentasProvider.On("Create", &mockVentaToCreate).Return(mockCarritoID, nil)
	mockCarritosProvider.On("Save", &mockEmptyCarrito).Return(mockCarritoID2, nil)
	mockUsuariosProvider.On("Save", &mockUpdatedUsuario).Return(nil)

	carritoID, err := useCase.Execute(command)

	assert.NoError(t, err)
	assert.Equal(t, carritoID, mockCarritoID2)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 2)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 1)
}

func TestWhenCarritoWasNotFoundShouldReturnNotFoundError(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarritosProvider.On("GetByID", command.CarritoID).
		Return(entities.Carrito{}, database.NewNotFoundError(errormessages.CarritoNotFound.String()))

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 0)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 0)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 0)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 0)
}

func TestWhenGetCarritoWentWrongShouldReturnInternalServerError(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarritosProvider.On("GetByID", command.CarritoID).
		Return(entities.Carrito{}, database.NewInternalServerError(errormessages.InternalServerError.String()))

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 0)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 0)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 0)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 0)
}

func TestWhenCarritoWasAlreadyPurchasedShouldReturnAlreadyPurchaseError(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockUpdatedCarrito := getMockUpdatedCarrito()

	mockCarritosProvider.On("GetByID", command.CarritoID).Return(mockUpdatedCarrito, nil)

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "business.CarritoAlreadyPurchaseError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 0)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 0)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 0)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 0)
}

func TestWhenUserWasNotFoundShouldReturnNotFoundError(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarrito := getMockCarrito()

	mockCarritosProvider.On("GetByID", command.CarritoID).
		Return(mockCarrito, nil)
	mockUsuariosProvider.On("GetByID", command.CompradorID).
		Return(entities.Usuario{}, database.NewNotFoundError(errormessages.UsuarioNotFound.String()))

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 0)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 0)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 0)
}

func TestWhenGetUserWentWrongShouldReturnInternalServerError(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarrito := getMockCarrito()

	mockCarritosProvider.On("GetByID", command.CarritoID).
		Return(mockCarrito, nil)
	mockUsuariosProvider.On("GetByID", command.CompradorID).
		Return(entities.Usuario{}, database.NewInternalServerError(errormessages.InternalServerError.String()))

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 0)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 0)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 0)
}

func TestWhenSaveUpdatedCarritoWentWrongThenShouldReturnInternalError(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarrito := getMockCarrito()
	mockUpdatedCarrito := getMockUpdatedCarrito()
	mockUsuario := getMockUsuario()

	mockCarritosProvider.On("GetByID", command.CarritoID).Return(mockCarrito, nil)
	mockUsuariosProvider.On("GetByID", command.CompradorID).Return(mockUsuario, nil)
	mockCarritosProvider.On("Save", &mockUpdatedCarrito).
		Return(uint64(0), database.NewInternalServerError(errormessages.InternalServerError.String()))

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 1)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 0)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 0)
}

func TestWhenCreateVentaWentWrongShouldReturnInternalError(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarrito := getMockCarrito()
	mockUpdatedCarrito := getMockUpdatedCarrito()
	mockVentaToCreate := getMockVentaToCreate()
	mockUsuario := getMockUsuario()

	mockCarritosProvider.On("GetByID", command.CarritoID).Return(mockCarrito, nil)
	mockUsuariosProvider.On("GetByID", command.CompradorID).Return(mockUsuario, nil)
	mockCarritosProvider.On("Save", &mockUpdatedCarrito).Return(uint64(0), nil)
	mockVentasProvider.On("Create", &mockVentaToCreate).
		Return(uint64(0), database.NewInternalServerError(errormessages.InternalServerError.String()))

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 1)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 0)
}

func TestWhenSaveEmptyCarritoWentWrongThenShouldReturn(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarrito := getMockCarrito()
	mockUpdatedCarrito := getMockUpdatedCarrito()
	mockVentaToCreate := getMockVentaToCreate()
	mockEmptyCarrito := getMockNewCarrito()
	mockUsuario := getMockUsuario()

	mockCarritosProvider.On("GetByID", command.CarritoID).Return(mockCarrito, nil)
	mockUsuariosProvider.On("GetByID", command.CompradorID).Return(mockUsuario, nil)
	mockCarritosProvider.On("Save", &mockUpdatedCarrito).Return(uint64(0), nil)
	mockVentasProvider.On("Create", &mockVentaToCreate).Return(mockCarritoID, nil)
	mockCarritosProvider.On("Save", &mockEmptyCarrito).
		Return(uint64(0), database.NewInternalServerError(errormessages.InternalServerError.String()))

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 2)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 0)
}

func TestWhenSaveUpdatedUserWentWrongThenShouldReturnInternalError(t *testing.T) {
	useCase := initialize()
	command := getMockCommand()

	mockCarrito := getMockCarrito()
	mockUpdatedCarrito := getMockUpdatedCarrito()
	mockVentaToCreate := getMockVentaToCreate()
	mockUsuario := getMockUsuario()
	mockUpdatedUsuario := getMockUpdatedUsuario()
	mockEmptyCarrito := getMockNewCarrito()

	mockCarritosProvider.On("GetByID", command.CarritoID).Return(mockCarrito, nil)
	mockUsuariosProvider.On("GetByID", command.CompradorID).Return(mockUsuario, nil)
	mockCarritosProvider.On("Save", &mockUpdatedCarrito).Return(uint64(0), nil)
	mockVentasProvider.On("Create", &mockVentaToCreate).Return(mockCarritoID, nil)
	mockCarritosProvider.On("Save", &mockEmptyCarrito).Return(mockCarritoID2, nil)
	mockUsuariosProvider.On("Save", &mockUpdatedUsuario).
		Return(database.NewInternalServerError(errormessages.InternalServerError.String()))

	carritoID, err := useCase.Execute(command)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
	assert.Empty(t, carritoID)
	mockCarritosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "GetByID", 1)
	mockCarritosProvider.AssertNumberOfCalls(t, "Save", 2)
	mockVentasProvider.AssertNumberOfCalls(t, "Create", 1)
	mockUsuariosProvider.AssertNumberOfCalls(t, "Save", 1)
}

func getMockCarrito() entities.Carrito {
	return entities.Carrito{
		ID:              mockCarritoID,
		Subtotal:        1000,
		Descuento:       5,
		Envio:           5,
		PrecioTotal:     1000,
		EstadoCarritoID: estadoscarrito.Active,
	}
}

func getMockUpdatedCarrito() entities.Carrito {
	return entities.Carrito{
		ID:              mockCarritoID,
		Subtotal:        1000,
		Descuento:       5,
		Envio:           5,
		PrecioTotal:     1000,
		EstadoCarritoID: estadoscarrito.Purchased,
	}
}

func getMockVentaToCreate() entities.Venta {
	return entities.Venta{
		MedioDePagoID: mediosdepago.Contraentrega,
		CarritoID:     getMockCarrito().ID,
		CompradorID:   getMockUsuario().ID,
	}
}

func getMockNewCarrito() entities.Carrito {
	return entities.Carrito{
		EstadoCarritoID: estadoscarrito.Active,
	}
}

func getMockUsuario() entities.Usuario {
	return entities.Usuario{
		ID:              2,
		Nombre:          "Frey",
		Apellido:        "Man",
		Password:        "MeQuieroMuch0",
		CarritoActualID: mockCarritoID,
	}
}

func getMockUpdatedUsuario() entities.Usuario {
	return entities.Usuario{
		ID:              2,
		Nombre:          "Frey",
		Apellido:        "Man",
		Password:        "MeQuieroMuch0",
		CarritoActualID: mockCarritoID2,
	}
}

func getMockCommand() command.PurchaseCarritoCommand {
	return command.PurchaseCarritoCommand{
		CarritoID:     mockCarritoID,
		CompradorID:   2,
		MedioDePagoID: mediosdepago.Contraentrega,
	}
}

func initialize() Implementation {
	mockCarritosProvider = new(mocks.MockCarritoProvider)
	mockUsuariosProvider = new(mocks.MockUsuarioProvider)
	mockVentasProvider = new(mocks.MockVentaProvider)
	return Implementation{
		CarritosProvider: mockCarritosProvider,
		VentasProvider:   mockVentasProvider,
		UsuariosProvider: mockUsuariosProvider,
	}
}
