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
	QueryUpdate     = "UPDATE `carritos` SET `subtotal`=?,`descuento`=?,`envio`=?,`precio_total`=?,`estado_carrito_id`=? WHERE `id` = ?"
	QueryCreate     = "INSERT INTO `carritos` (`subtotal`,`descuento`,`envio`,`precio_total`,`estado_carrito_id`) VALUES (?,?,?,?,?)"
)

func TestGetBycodeOK(t *testing.T) {
	initialize()

	carrito := getResponse()

	mockDB.ExpectQuery(QuerySelectByID).WithArgs(carrito.ID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "precio_total", "descuento", "subtotal", "envio"}).AddRow(
			carrito.ID, carrito.PrecioTotal, carrito.Descuento, carrito.Subtotal, carrito.Envio,
		),
	)
	carritoRecibido, err := repository.GetByID(carrito.ID)

	assert.NoError(t, err)
	assert.Equal(t, carrito, carritoRecibido)
}

func TestGetByCodeErrorNotFound(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrRecordNotFound)

	carritoRecibido, err := repository.GetByID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, carritoRecibido)
}

func TestGetByCodeInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetByID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestWhenSaveWasSuccesfullShouldReturnNoError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdate).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()
	mockCarrito := getMockCarrito()

	err := repository.Save(&mockCarrito)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenSaveAndCarritoDoesNotContainIDShouldCreateAndReturnNoError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryCreate).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()
	mockCarrito := getMockCarritoToCreate()

	err := repository.Save(&mockCarrito)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenSaveWentWrongShouldReturnInternalError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdate).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectBegin()
	mockCarrito := getMockCarrito()

	err := repository.Save(&mockCarrito)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getResponse() entities.Carrito {
	return entities.Carrito{
		ID:          mockCarritoID,
		PrecioTotal: 1000,
		Subtotal:    995,
		Envio:       5,
		Descuento:   0,
	}
}

func getMockCarrito() entities.Carrito {
	return entities.Carrito{
		ID:          1,
		Subtotal:    1000,
		Descuento:   5,
		Envio:       5,
		PrecioTotal: 1000,
	}
}

func getMockCarritoToCreate() entities.Carrito {
	return entities.Carrito{
		Subtotal:    1000,
		Descuento:   5,
		Envio:       5,
		PrecioTotal: 1000,
	}
}
