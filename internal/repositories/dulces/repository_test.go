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
	MockDulceID       = uint64(132423)
	QuerySelectByCode = "Call GetDetalleDulceByCode(?)"
	QuerySelectByID   = "Call GetDetalleDulceByID(?)"
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
