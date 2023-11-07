package dulces

import (
	"reflect"
	"testing"
	"time"

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
	mockDulceID1                     = uint64(1231)
	mockDulceID2                     = uint64(2321)
	mockCarritoID                    = uint64(2321)
	QuerySelectByCode                = "Call GetDetalleDulceByCode(?)"
	QuerySelectByID                  = "Call GetDetalleDulceByID(?)"
	QuerySelectDulcesListByCarritoID = "SELECT * FROM `carritos_dulces` WHERE carrito_id = ?"
	QueryGetByID                     = "SELECT * FROM `dulces` WHERE id = ? LIMIT 1"
)

func TestGetBycodeOK(t *testing.T) {
	initialize()

	dulce := getResponse()

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
	initialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2").WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetByCode("2")

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetByCodeInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2").WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetByCode("2")

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDetailByIDOK(t *testing.T) {
	initialize()

	dulce := getResponse()

	mockDB.ExpectQuery(QuerySelectByID).WithArgs(dulce.ID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre", "presentacion_id", "presentacion_nombre", "descripcion", "imagen", "disponibles", "precio_unidad", "peso", "marca_id", "marca_nombre", "codigo"}).AddRow(
			dulce.ID, dulce.Nombre, dulce.Presentacion.ID, dulce.Presentacion.Nombre, dulce.Descripcion, dulce.Imagen, dulce.Disponibles, dulce.PrecioUnidad, dulce.Peso, dulce.Marca.ID, dulce.Marca.Nombre, dulce.Codigo,
		),
	)
	dulceRecibido, err := repository.GetDetailByID(dulce.ID)

	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
}

func TestGetDetailByIDInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(MockDulceID).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetDetailByID(MockDulceID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDetailByIDErrorNotFound(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(MockDulceID).WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetDetailByID(MockDulceID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetDulcesListByCarritoIDWhenEveryThingWentSuccessfullyShouldReturnDulcesList(t *testing.T) {
	initialize()
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
	initialize()

	mockDB.ExpectQuery(QuerySelectDulcesListByCarritoID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrInvalidData)

	dulcesListResponse, err := repository.GetDulcesListByCarritoID(mockCarritoID)

	assert.Error(t, err)
	assert.Empty(t, dulcesListResponse)
	assert.NoError(t, mockDB.ExpectationsWereMet())
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

func TestGetByIDOK(t *testing.T) {
	initialize()

	dulce := getMockDulce()

	mockDB.ExpectQuery(QueryGetByID).WithArgs(dulce.ID).WillReturnRows(
		sqlmock.NewRows([]string{
			"id", "nombre", "marca_id", "precio_unidad", "peso", "unidades", "presentacion_id", "descripcion",
			"imagen", "fecha_vencimiento", "fecha_expedicion", "disponibles", "codigo",
		}).AddRow(
			dulce.ID, dulce.Nombre, dulce.MarcaID, dulce.PrecioUnidad, dulce.Peso, dulce.Unidades, dulce.PresentacionID, dulce.Descripcion,
			dulce.Imagen, dulce.FechaVencimiento, dulce.FechaExpedicion, dulce.Disponibles, dulce.Codigo,
		),
	)
	dulceRecibido, err := repository.GetByID(dulce.ID)

	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
}

func TestGetByIDErrorNotFound(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QueryGetByID).WithArgs(2).WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetByID(2)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetByIDInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QueryGetByID).WithArgs(2).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetByID(2)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getResponse() (response query.DetalleDulce) {
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

func getMockDulce() (dulce entities.Dulce) {
	dulce = entities.Dulce{
		ID:               1,
		Nombre:           "Gomas Clasicas",
		MarcaID:          6,
		PrecioUnidad:     2950.000,
		Peso:             80,
		Unidades:         5,
		PresentacionID:   4,
		Descripcion:      "Gomas clasicas con sobores surtidos",
		FechaVencimiento: time.Date(2023, time.August, 24, 0, 0, 0, 0, time.Local),
		FechaExpedicion:  time.Date(2023, time.July, 24, 0, 0, 0, 0, time.Local),
		Disponibles:      100,
		Codigo:           "1A",
	}
	return
}
