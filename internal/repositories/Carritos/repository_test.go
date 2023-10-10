package carritos

import (
	dbmocks "Mileyman-API/internal/app/config/database/mocks"
	"Mileyman-API/internal/domain/entities"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	repository Repository
	mockDB     sqlmock.Sqlmock
	DB         *gorm.DB
)

const (
	QueryGetDulceByCarritoIDAndDulceID = "SELECT * FROM `carritos_dulces` WHERE carrito_id = ? AND dulce_id = ? ORDER BY `carritos_dulces`.`id` LIMIT 1"
	QueryUpdateDulceInCarrito          = ""
	QueryAddDulceInCarrito             = ""
)

func TestGetDulceByCarritoIDAndDulceIDOK(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulce()

	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(carritoDulce.CarritoID, carritoDulce.DulceID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "carrito_id", "dulce_id", "unidades", "subtotal"}).AddRow(
			carritoDulce.ID, carritoDulce.CarritoID, carritoDulce.DulceID, carritoDulce.Unidades, carritoDulce.Subtotal,
		),
	)
	carritoDulceRecibido, exist, err := repository.GetDulceByCarritoIDAndDulceID(carritoDulce.CarritoID, carritoDulce.DulceID)

	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, carritoDulce, carritoDulceRecibido)

}

func TestGetDulceByCarritoIDAndDulceIDWhenCartDoesNotExistReturnsFalse(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(2, 2).WillReturnError(gorm.ErrRecordNotFound)

	carritoDulceRecibido, exist, err := repository.GetDulceByCarritoIDAndDulceID(2, 2)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.False(t, exist)
	assert.Empty(t, carritoDulceRecibido)

}

func TestUpdateDulceInCarritoOK(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
}

func inicialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
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
