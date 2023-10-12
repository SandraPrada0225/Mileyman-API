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
	MockDulceID       = uint64(132423)
	QuerySelectByCode = "Call GetDetalleDulceByCode(?)"
	QuerySelectByID   = "Call GetDetalleDulceByID(?)"
	QueryGetByID = "SELECT * FROM `dulces` WHERE id = ? LIMIT 1"
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

func TestByCodeErrorNotFound(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs("2").WillReturnError(gorm.ErrRecordNotFound)

	dulceRecibido, err := repository.GetByCode("2")

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestByCodeInternalServerError(t *testing.T) {
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

func TestGetByIDOK(t *testing.T) {
	inicialize()

	dulce := GetMockDulce()

	mockDB.ExpectQuery(QueryGetByID).WithArgs(dulce.ID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre", "marca_id", "precio_unidad", "peso", "unidades", "presentacion_id", "descripcion",
		 "imagen", "fecha_vencimiento", "fecha_expedicion", "disponibles", "codigo"}).AddRow(
			dulce.ID, dulce.Nombre, dulce.MarcaID, dulce.PrecioUnidad, dulce.Peso, dulce.Unidades, dulce.PresentacionID, dulce.Descripcion, 
			dulce.Imagen, dulce.FechaVencimiento, dulce.FechaExpedicion, dulce.Disponibles, dulce.Codigo,
		),
	)
	dulceRecibido, err := repository.GetByID(dulce.ID)

	assert.NoError(t, err)
	assert.Equal(t, dulce, dulceRecibido)
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

func GetMockDulce() (dulce entities.Dulce) {
	dulce = entities.Dulce{
		ID: 1,
		Nombre: "Gomas Clasicas",
		MarcaID: 6,
		PrecioUnidad: 2950.000,
		Peso: 80,
		Unidades: 5,
		PresentacionID: 4,
		Descripcion: "Gomas clasicas con sobores surtidos",
		FechaVencimiento: time.Date(2023, time.August, 24, 0, 0, 0, 0, time.Local),
		FechaExpedicion: time.Date(2023, time.July, 24, 0, 0, 0, 0, time.Local),
		Disponibles: 100,
		Codigo: "1A",
	}
	return
}
