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
	mockCarritoID                      = uint64(2)
	QuerySelectByID                    = "SELECT * FROM `carritos` WHERE id = ? LIMIT 1"
	QueryGetDulceByCarritoIDAndDulceID = "SELECT * FROM `carritos_dulces` WHERE carrito_id = ? AND dulce_id = ? ORDER BY `carritos_dulces`.`id` LIMIT 1"
	QueryUpdateDulceInCarrito          = "UPDATE `carritos_dulces` SET `carrito_id`=?,`dulce_id`=?,`unidades`=?,`subtotal`=? WHERE `id` = ?"
	QueryAddDulceInCarrito             = "INSERT INTO `carritos_dulces` (`carrito_id`,`dulce_id`,`unidades`,`subtotal`) VALUES (?,?,?,?)"
	QueryDeleteDulceInCarrito          = "DELETE FROM `carritos_dulces` WHERE `carritos_dulces`.`id` = ?"
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

func TestByCodeErrorNotFound(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrRecordNotFound)

	carritoRecibido, err := repository.GetCarritoByCarritoID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, carritoRecibido)
}

func TestByCodeInternalServerError(t *testing.T) {
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

	assert.NoError(t, err)
	assert.False(t, exist)
	assert.Empty(t, carritoDulceRecibido)

}

func TestGetDulceByCarritoIDAndDulceIDInternalServerError(t *testing.T) {
	inicialize()

	mockDB.ExpectBegin()
	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(2, 2).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	carritoDulceRecibido, exist, err := repository.GetDulceByCarritoIDAndDulceID(2, 2)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.False(t, exist)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, carritoDulceRecibido)
}


func TestUpdateDulceInCarritoOK(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.AddDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
}
func TestUpdateDulceInCarritoNotFoundError(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
}

func TestUpdateDulceInCarritoInternalServerError(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.DeleteDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
}

func TestAddDulceInCarritoOK(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.AddDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
}

func TestAddDulceInCarritoNotFoundError(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
}
func TestAddDulceInCarritoInternalServerError(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
}
func TestDeleteDulceInCarritoOK(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.DeleteDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
}

func TestDeleteDulceInCarritoNotFoundError(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.DeleteDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
}

func TestDeleteDulceInCarritoInternalServerError(t *testing.T) {
	inicialize()

	carritoDulce := GetMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.DeleteDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
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

func GetMockCarritoDulceSinID() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{

		CarritoID: 1,
		DulceID:   1,
		Unidades:  2,
		Subtotal:  5900,
	}
	return
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
