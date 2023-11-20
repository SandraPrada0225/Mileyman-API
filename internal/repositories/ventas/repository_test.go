package ventas

import (
	"reflect"
	"testing"
	"time"

	dbmocks "Mileyman-API/internal/app/config/database/mocks"
	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	createQuery                 = "INSERT INTO `ventas` (`medio_de_pago_id`,`carrito_id`,`comprador_id`,`created_at`) VALUES (?,?,?,?)"
	queryGetListByUserID        = "Call GetPurchaseListByUserID(?)"
	mockUserID           uint64 = 4213
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

func TestGetPurchaseListWhenIsSuccesfullShouldReturnList(t *testing.T) {
	initialize()

	purchaseList := getMockedPurchaseList()

	mockDB.ExpectQuery(queryGetListByUserID).WithArgs(mockUserID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "fecha", "medio_de_pago_id", "medio_de_pago", "carrito_id", "precio_total", "subtotal", "descuento", "envio", "estado_carrito_id", "estado_carrito"}).
			AddRow(
				purchaseList.PurchaseList[0].ID, purchaseList.PurchaseList[0].Fecha, purchaseList.PurchaseList[0].MedioDePagoID, purchaseList.PurchaseList[0].MedioDePago, purchaseList.PurchaseList[0].CarritoID,
				purchaseList.PurchaseList[0].PrecioTotal, purchaseList.PurchaseList[0].Subtotal, purchaseList.PurchaseList[0].Descuento, purchaseList.PurchaseList[0].Envio,
				purchaseList.PurchaseList[0].EstadoCarritoID, purchaseList.PurchaseList[0].EstadoCarrito).
			AddRow(
				purchaseList.PurchaseList[1].ID, purchaseList.PurchaseList[1].Fecha, purchaseList.PurchaseList[0].MedioDePagoID, purchaseList.PurchaseList[1].MedioDePago, purchaseList.PurchaseList[1].CarritoID,
				purchaseList.PurchaseList[1].PrecioTotal, purchaseList.PurchaseList[1].Subtotal, purchaseList.PurchaseList[1].Descuento, purchaseList.PurchaseList[1].Envio,
				purchaseList.PurchaseList[1].EstadoCarritoID, purchaseList.PurchaseList[1].EstadoCarrito),
	)

	response, err := repository.GetListByUserID(mockUserID)

	assert.NoError(t, err)
	assert.Equal(t, purchaseList, response)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetPurchaseListWhenWentWrongShouldReturnInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(queryGetListByUserID).WithArgs(mockUserID).WillReturnError(gorm.ErrInvalidData)

	response, err := repository.GetListByUserID(mockUserID)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
	assert.Empty(t, response)
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

func getMockedPurchaseList() responses.GetPurchaseList {
	fecha := time.Date(2023, 10, 17, 0, 0, 0, 0, time.Local)
	return responses.GetPurchaseList{
		PurchaseList: []responses.Purchase{
			{
				ID:              1,
				Fecha:           fecha,
				MedioDePagoID:   1,
				MedioDePago:     "contraentrega",
				PrecioTotal:     100,
				Subtotal:        97,
				Descuento:       2,
				Envio:           5,
				EstadoCarritoID: 1,
				EstadoCarrito:   "comprado",
			},
			{
				ID:              2,
				Fecha:           fecha,
				MedioDePagoID:   1,
				MedioDePago:     "credito",
				PrecioTotal:     120,
				Subtotal:        100,
				Descuento:       0,
				Envio:           20,
				EstadoCarritoID: 1,
				EstadoCarrito:   "comprado",
			},
		},
	}
}
