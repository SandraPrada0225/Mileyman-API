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
	querySelectByID                    = "SELECT * FROM `carritos` WHERE id = ? LIMIT 1"
	queryUpdate                        = "UPDATE `carritos` SET `subtotal`=?,`descuento`=?,`envio`=?,`precio_total`=?,`estado_carrito_id`=? WHERE `id` = ?"
	queryCreate                        = "INSERT INTO `carritos` (`subtotal`,`descuento`,`envio`,`precio_total`,`estado_carrito_id`) VALUES (?,?,?,?,?)"
	queryGetDulceByCarritoIDAndDulceID = "SELECT * FROM `carritos_dulces` WHERE carrito_id = ? AND dulce_id = ? ORDER BY `carritos_dulces`.`id` LIMIT 1"
	queryUpdateDulceInCarrito          = "UPDATE `carritos_dulces` SET `carrito_id`=?,`dulce_id`=?,`unidades`=?,`subtotal`=? WHERE `id` = ?"
	queryAddDulceInCarrito             = "INSERT INTO `carritos_dulces` (`carrito_id`,`dulce_id`,`unidades`,`subtotal`) VALUES (?,?,?,?)"
	queryDeleteDulceInCarrito          = "DELETE FROM `carritos_dulces` WHERE `carritos_dulces`.`id` = ?"
)

func TestGetBycodeOK(t *testing.T) {
	initialize()

	carrito := getResponse()

	mockDB.ExpectQuery(querySelectByID).WithArgs(carrito.ID).WillReturnRows(
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
	mockDB.ExpectQuery(querySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrRecordNotFound)

	carritoRecibido, err := repository.GetByID(mockCarritoID)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", typeErr)
	assert.Empty(t, carritoRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetByCodeInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(querySelectByID).WithArgs(mockCarritoID).WillReturnError(gorm.ErrInvalidData)

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
	mockDB.ExpectExec(queryUpdate).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()
	mockCarrito := getMockCarrito()

	err := repository.Save(&mockCarrito)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenSaveAndCarritoDoesNotContainIDShouldCreateAndReturnNoError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(queryCreate).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()
	mockCarrito := getMockCarritoToCreate()

	err := repository.Save(&mockCarrito)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestWhenSaveWentWrongShouldReturnInternalError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(queryUpdate).WillReturnError(gorm.ErrInvalidData)
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

	mockDB.ExpectQuery(queryGetDulceByCarritoIDAndDulceID).WithArgs(carritoDulce.CarritoID, carritoDulce.DulceID).WillReturnRows(
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
	mockDB.ExpectQuery(queryGetDulceByCarritoIDAndDulceID).WithArgs(2, 2).WillReturnError(gorm.ErrRecordNotFound)

	carritoDulceRecibido, exist, err := repository.GetDulceByCarritoIDAndDulceID(2, 2)

	assert.NoError(t, err)
	assert.False(t, exist)
	assert.Empty(t, carritoDulceRecibido)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetDulceByCarritoIDAndDulceIDInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(queryGetDulceByCarritoIDAndDulceID).WithArgs(2, 2).WillReturnError(gorm.ErrInvalidData)

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
	mockDB.ExpectExec(queryUpdateDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.AddDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestUpdateDulceInCarritoNotFoundError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(queryUpdateDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
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
	mockDB.ExpectExec(queryUpdateDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
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
	mockDB.ExpectExec(queryAddDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.AddDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestAddDulceInCarritoNotFoundError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulceSinID()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(queryAddDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
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
	mockDB.ExpectExec(queryAddDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
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
	mockDB.ExpectExec(queryDeleteDulceInCarrito).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	err := repository.DeleteDulceInCarrito(carritoDulce)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestDeleteDulceInCarritoNotFoundError(t *testing.T) {
	initialize()

	carritoDulce := getMockCarritoDulce()

	mockDB.ExpectBegin()
	mockDB.ExpectExec(queryDeleteDulceInCarrito).WillReturnError(gorm.ErrRecordNotFound)
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
	mockDB.ExpectExec(queryDeleteDulceInCarrito).WillReturnError(gorm.ErrInvalidData)
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
