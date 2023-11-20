package getpurchaselistbyuserid

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/usecase/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockGetPurchaseListByUserID *mocks.MockGetPurchaseListByUSerID

const (
	mockUserID uint64 = 5
)

func CreateServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	mockGetPurchaseListByUserID = new(mocks.MockGetPurchaseListByUSerID)
	handler := GetPurchaseListByUserID{
		UseCase: mockGetPurchaseListByUserID,
	}
	r := gin.Default()
	group := r.Group("/api/users")
	group.GET("/:id/purchases", handler.Handle())
	return r
}

func TestOk(t *testing.T) {
	r := CreateServer()
	expectedResponse := getMockedPurchaseList()

	json, _ := json.Marshal(&expectedResponse)

	mockGetPurchaseListByUserID.On("Execute", mockUserID).Return(expectedResponse, nil)
	request := httptest.NewRequest("GET", "/api/users/5/purchases", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	bodyResponse, err := io.ReadAll(response.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(json), string(bodyResponse))
	mockGetPurchaseListByUserID.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenWentWrongThenShouldReturn500InternalError(t *testing.T) {
	r := CreateServer()

	mockGetPurchaseListByUserID.On("Execute", mockUserID).Return(entities.Dulce{}, database.NewInternalServerError("error"))
	request := httptest.NewRequest("GET", "/api/users/5/purchases", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockGetPurchaseListByUserID.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenIDIsNotValidThenShouldReturn400BadRequest(t *testing.T) {
	r := CreateServer()

	mockGetPurchaseListByUserID.On("Execute", mockUserID).Return(entities.Dulce{}, database.NewInternalServerError("error"))
	request := httptest.NewRequest("GET", "/api/users/5a/purchases", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	mockGetPurchaseListByUserID.AssertNumberOfCalls(t, "Execute", 0)
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
