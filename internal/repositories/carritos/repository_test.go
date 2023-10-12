package carritos

import (
	"reflect"
	"testing"

	dbmocks "Mileyman-API/internal/app/config/database/mocks"
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
	mockCarritoID   = uint64(2)
	QuerySelectByID = "SELECT * FROM `carritos` WHERE id = ? LIMIT 1"
)

func TestGetBycodeOK(t *testing.T) {
	inicialize()

	carrito := GetResponse()

	mockDB.ExpectQuery(QuerySelectByID).WithArgs(carrito.ID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "precio_total", "descuento", "sub_total", "envio"}).AddRow(
			carrito.ID, carrito.PrecioTotal, carrito.Descuento, carrito.SubTotal, carrito.Envio,
		),
	)
	carritoRecibido, err := repository.GetCarritoByCarritoID(carrito.ID)

	assert.NoError(t, err)
	assert.Equal(t, carrito, carritoRecibido)
}

func TestGetByCodeErrorNotFound(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrRecordNotFound)

	carritoRecibido, err := repository.GetCarritoByCarritoID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, carritoRecibido)
}

func TestGetByCodeInternalServerError(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetCarritoByCarritoID(mockCarritoID)

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

func GetResponse() (response entities.Carrito) {
	response = entities.Carrito{
		ID:          mockCarritoID,
		PrecioTotal: 1000,
		SubTotal:    995,
		Envio:       5,
		Descuento:   0,
	}

	return
}
