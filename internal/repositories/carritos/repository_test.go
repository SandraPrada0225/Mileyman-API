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
	QueryUpdate                        = "UPDATE `carritos` SET `subtotal`=?,`descuento`=?,`envio`=?,`precio_total`=?,`estado_carrito_id`=? WHERE `id` = ?"
	QueryCreate                        = "INSERT INTO `carritos` (`subtotal`,`descuento`,`envio`,`precio_total`,`estado_carrito_id`) VALUES (?,?,?,?,?)"
	QueryGetDulceByCarritoIDAndDulceID = "SELECT * FROM `carritos_dulces` WHERE carrito_id = ? AND dulce_id = ? ORDER BY `carritos_dulces`.`id` LIMIT 1"
	QueryUpdateDulceInCarrito          = "UPDATE `carritos_dulces` SET `carrito_id`=?,`dulce_id`=?,`unidades`=?,`subtotal`=? WHERE `id` = ?"
	QueryAddDulceInCarrito             = "INSERT INTO `carritos_dulces` (`carrito_id`,`dulce_id`,`unidades`,`subtotal`) VALUES (?,?,?,?)"
	QueryDeleteDulceInCarrito          = "DELETE FROM `carritos_dulces` WHERE `carritos_dulces`.`id` = ?"
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
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetByCodeErrorNotFound(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrRecordNotFound)

	carritoRecibido, err := repository.GetByID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, carritoRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetByCodeInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetByID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
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

func TestGetDulceByCarritoIDAndDulceIDOK(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(carritoDulce.CarritoID, carritoDulce.DulceID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "carrito_id", "dulce_id", "unidades", "subtotal"}).AddRow(
			carritoDulce.ID, carritoDulce.CarritoID, carritoDulce.DulceID, carritoDulce.Unidades, carritoDulce.Subtotal,
		),
	)
	carritoDulceRecibido, exist, err := repository.GetDulceByCarritoIDAndDulceID(carritoDulce.CarritoID, carritoDulce.DulceID)

	assert.NoError(t, err)
	assert.True(t, exist)
	assert.Equal(t, carritoDulce, carritoDulceRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetDulceByCarritoIDAndDulceIDWhenCartDoesNotExistReturnsFalse(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(2, 2).WillReturnError(gorm.ErrRecordNotFound)

	carritoDulceRecibido, exist, err := repository.GetDulceByCarritoIDAndDulceID(2, 2)

	assert.NoError(t, err)
	assert.False(t, exist)
	assert.Empty(t, carritoDulceRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetDulceByCarritoIDAndDulceIDInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QueryGetDulceByCarritoIDAndDulceID).WithArgs(2, 2).WillReturnError(gorm.ErrInvalidData)

	carritoDulceRecibido, exist, err := repository.GetDulceByCarritoIDAndDulceID(2, 2)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.False(t, exist)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, carritoDulceRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestUpdateDulceInCarritoOK(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.AddDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestUpdateDulceInCarritoNotFoundError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestUpdateDulceInCarritoInternalServerError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryUpdateDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.DeleteDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
}

func TestAddDulceInCarritoOK(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.AddDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestAddDulceInCarritoNotFoundError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestAddDulceInCarritoInternalServerError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryAddDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.AddDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestDeleteDulceInCarritoOK(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.DeleteDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestDeleteDulceInCarritoNotFoundError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
	mockDB.ExpectRollback()

	err := repository.DeleteDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestDeleteDulceInCarritoInternalServerError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(QueryDeleteDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()

	err := repository.DeleteDulceInCarrito(carritoDulce)
	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func getMockCarritoDulce() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		ID:        1,
		CarritoID: 1,
		DulceID:   1,
		Unidades:  2,
		Subtotal:  5900,
	}
	return
}

func getMockCarritoDulceSinID() (carritoDulce entities.CarritoDulce) {
	carritoDulce = entities.CarritoDulce{
		CarritoID: 1,
		DulceID:   1,
		Unidades:  2,
		Subtotal:  5900,
	}
	return
}

func getResponse() (response entities.Carrito) {
	response = entities.Carrito{
		ID:          mockCarritoID,
		PrecioTotal: 1000,
		Subtotal:    995,
		Envio:       5,
		Descuento:   0,
	}
	return
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
