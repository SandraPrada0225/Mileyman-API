package dulces

import (
	"reflect"
	"testing"

	dbmocks "Mileyman-API/internal/app/config/database/mocks"
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

var (
	repository Repository
	mockDB     sqlmock.Sqlmock
	DB         *gorm.DB
)

const (
	MockDulceID                      = uint64(132423)
	mockCarritoID                    = uint64(2321)
	mockDulceID1                     = uint64(1231)
	mockDulceID2                     = uint64(2321)
	QuerySelectByCode                = "Call GetDetalleDulceByCode(?)"
	QuerySelectByID                  = "Call GetDetalleDulceByID(?)"
	QuerySelectDulcesListByCarritoID = "SELECT * FROM `carritos_dulces` WHERE carrito_id = ?"
)

func TestGetBycodeOK(t *testing.T) {
	inicialize()

	dulce := GetResponse()

	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(dulce.Codigo).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre", "presentacion_id", "presentacion_nombre", "descripcion", "imagen", "disponibles", "precio_unidad", "peso", "marca_id", "marca_nombre", "codigo"}).AddRow(
			dulce.ID, dulce.Nombre, dulce.Presentacion.ID, dulce.Presentacion.Nombre, dulce.Descripcion, dulce.Imagen, dulce.Disponibles, dulce.PrecioUnidad, dulce.Peso, dulce.Marca.ID, dulce.Marca.Nombre, dulce.Codigo,
		),
	)
	dulceRecibido, err := repository.GetByCode(dulce.Codigo)

	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
}

func TestGetByCodeErrorNotFound(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2").WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetByCode("2")

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetByCodeInternalServerError(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2").WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetByCode("2")

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDetailByIDOK(t *testing.T) {
	inicialize()

	dulce := GetResponse()

	mockDB.ExpectQuery(QuerySelectByID).WithArgs(dulce.ID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre", "presentacion_id", "presentacion_nombre", "descripcion", "imagen", "disponibles", "precio_unidad", "peso", "marca_id", "marca_nombre", "codigo"}).AddRow(
			dulce.ID, dulce.Nombre, dulce.Presentacion.ID, dulce.Presentacion.Nombre, dulce.Descripcion, dulce.Imagen, dulce.Disponibles, dulce.PrecioUnidad, dulce.Peso, dulce.Marca.ID, dulce.Marca.Nombre, dulce.Codigo,
		),
	)
	dulceRecibido, err := repository.GetDetailByID(dulce.ID)

	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
}

func TestGetDetailByIDErrorNotFound(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs().WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetDetailByID(MockDulceID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDetailByIDInternalServerError(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(MockDulceID).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetDetailByID(MockDulceID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDulcesListByCarritoIDWhenEveryThingWentSuccessfullyShouldReturnDulcesList(t *testing.T) {
	inicialize()
	dulcesList := getMockDulcesInCarritoList()

	mockDB.ExpectQuery(QuerySelectDulcesListByCarritoID).WithArgs(mockCarritoID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "dulce_id", "carrito_id", "unidades", "subtotal"}).
			AddRow(dulcesList[0].ID, dulcesList[0].DulceID, dulcesList[0].CarritoID, dulcesList[0].Unidades, dulcesList[0].Subtotal).
			AddRow(dulcesList[1].ID, dulcesList[1].DulceID, dulcesList[1].CarritoID, dulcesList[1].Unidades, dulcesList[1].Subtotal),
	)

	dulcesListResponse, err := repository.GetDulcesListByCarritoID(mockCarritoID)

	assert.NoError(t, err)
	assert.Equal(t, dulcesList, dulcesListResponse)
}

func TestGetDulcesListByCarritoIDWhenSomethingWentWrongShouldReturnInternalError(t *testing.T) {
	inicialize()

	mockDB.ExpectQuery(QuerySelectDulcesListByCarritoID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrInvalidData)

	dulcesListResponse, err := repository.GetDulcesListByCarritoID(mockCarritoID)

	assert.Error(t, err)
	assert.Empty(t, dulcesListResponse)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func inicialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func GetResponse() (response query.DetalleDulce) {
	response = query.DetalleDulce{
		ID:     MockDulceID,
		Nombre: "Chocolatina",
		Presentacion: entities.Presentacion{
			ID:     1,
			Nombre: "Empaque",
		},
		Descripcion:  "Deliciosa chocolatina que se derrite en tu boca",
		Imagen:       "imagen",
		Disponibles:  100,
		PrecioUnidad: 1000,
		Peso:         40,
		Marca: entities.Marca{
			ID:     1,
			Nombre: "Jet",
		},
		Codigo: "2",
	}
	return
}

func getMockDulcesInCarritoList() []entities.CarritoDulce {
	return []entities.CarritoDulce{
		{
			ID:        1,
			DulceID:   1,
			CarritoID: 1,
			Unidades:  2,
			Subtotal:  2000,
		},
		{
			ID:        2,
			DulceID:   2,
			CarritoID: 1,
			Unidades:  1,
			Subtotal:  1000,
		},
	}
}
