package ventas

import (
	"reflect"
	"testing"
	"time"

	dbmocks "Mileyman-API/internal/app/config/database/mocks"
	"Mileyman-API/internal/domain/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	createQuery = "INSERT INTO `ventas` (`medio_de_pago_id`,`carrito_id`,`comprador_id`,`created_at`) VALUES (?,?,?,?)"
)

var (
	repository Repository
	mockDB     sqlmock.Sqlmock
	DB         *gorm.DB
)

func TestWhenCreatedWasSuccesfullyShouldReturnNoError(t *testing.T) {
	initialize()

	VentaToCreate := getMockedVentaToCreate()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(createQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.Create(&VentaToCreate)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenCreatedWhenWrongShouldReturnInternalError(t *testing.T) {
	initialize()

	VentaToCreate := getMockedVentaToCreate()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(createQuery).WillReturnError(gorm.ErrInvalidDB)
	mockDB.ExpectRollback()

	err := repository.Create(&VentaToCreate)
	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", errType)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenCarritoWasNotFoundShouldReturnNotFoundError(t *testing.T) {
	initialize()

	VentaToCreate := getMockedVentaToCreate()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(createQuery).WillReturnError(gorm.ErrForeignKeyViolated)
	mockDB.ExpectRollback()

	err := repository.Create(&VentaToCreate)
	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", errType)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenCarritoWasPurchasedShouldReturnConflictError(t *testing.T) {
	initialize()

	VentaToCreate := getMockedVentaToCreate()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(createQuery).WillReturnError(gorm.ErrDuplicatedKey)
	mockDB.ExpectRollback()

	err := repository.Create(&VentaToCreate)
	errType := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.ConflictError", errType)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getMockedVentaToCreate() entities.Venta {
	fecha := time.Date(2023, 10, 17, 0, 0, 0, 0, time.Local)
	return entities.Venta{
		CarritoID:     1,
		CreatedAt:     fecha,
		MedioDePagoID: 1,
	}
}
