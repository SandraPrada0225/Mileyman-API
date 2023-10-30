package purchasecarrito

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"Mileyman-API/cmd/server/handlers/purchase_carrito/contract"
	"Mileyman-API/cmd/server/handlers/purchase_carrito/response"
	"Mileyman-API/internal/domain/dto/command"
	"Mileyman-API/internal/domain/errors/business"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"
	"Mileyman-API/internal/use_case/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockPurchaseCarrito *mocks.MockPurchaseCarrito

const (
	mockMedioDePagoID uint64 = 1
	mockCompradorID   uint64 = 2
	mockNewCarritoID  uint64 = 3
)

func TestWhenRequestWasProcessedSuccesfullyThenShouldReturn200OkAndNewCarrito(t *testing.T) {
	r := createServer()
	mockCommand := getMockCommand()
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(getMockBodySuccess())
	if err != nil {
		panic(err.Error())
	}

	mockPurchaseCarrito.On("Execute", mockCommand).Return(mockNewCarritoID, nil)

	request := httptest.NewRequest("PUT", "/api/carritos/3/comprar", &body)
	request.Header.Add("Content-type", "application/json")

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	bodyResponse, err := io.ReadAll(response.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(getMockExpectedResponse()), string(bodyResponse))
	mockPurchaseCarrito.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenCarritoIDIsInvalidShouldReturBadRequest(t *testing.T) {
	r := createServer()
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(getMockBodySuccess())
	if err != nil {
		panic(err.Error())
	}

	request := httptest.NewRequest("PUT", "/api/carritos/3a/comprar", &body)
	request.Header.Add("Content-type", "application/json")

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	mockPurchaseCarrito.AssertNumberOfCalls(t, "Execute", 0)
}

func TestWhenBodyIsInvalidShouldReturnBadRequest(t *testing.T) {
	r := createServer()
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode("{bad body}")
	if err != nil {
		panic(err.Error())
	}

	request := httptest.NewRequest("PUT", "/api/carritos/3/comprar", &body)
	request.Header.Add("Content-type", "application/json")

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	mockPurchaseCarrito.AssertNumberOfCalls(t, "Execute", 0)
}

func TestWhenUseCaseWentWrongShouldReturnInternalError(t *testing.T) {
	r := createServer()
	mockCommand := getMockCommand()
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(getMockBodySuccess())
	if err != nil {
		panic(err.Error())
	}

	mockPurchaseCarrito.On("Execute", mockCommand).
		Return(uint64(0), database.NewInternalServerError(errormessages.InternalServerError.String()))

	request := httptest.NewRequest("PUT", "/api/carritos/3/comprar", &body)
	request.Header.Add("Content-type", "application/json")

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockPurchaseCarrito.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenCarritoNotFoundShouldReturnNotFoundError(t *testing.T) {
	r := createServer()
	mockCommand := getMockCommand()
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(getMockBodySuccess())
	if err != nil {
		panic(err.Error())
	}

	mockPurchaseCarrito.On("Execute", mockCommand).
		Return(uint64(0), database.NewNotFoundError(errormessages.CarritoNotFound.String()))

	request := httptest.NewRequest("PUT", "/api/carritos/3/comprar", &body)
	request.Header.Add("Content-type", "application/json")

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, response.Code)
	mockPurchaseCarrito.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenCarritoWasAlreadyPurchasedShouldReturnPreconditionFailed(t *testing.T) {
	r := createServer()
	mockCommand := getMockCommand()
	var body bytes.Buffer

	err := json.NewEncoder(&body).Encode(getMockBodySuccess())
	if err != nil {
		panic(err.Error())
	}

	mockPurchaseCarrito.On("Execute", mockCommand).
		Return(uint64(0), business.NewCarritoAlreadyPurchasedError(errormessages.CarritoHasAlreadyBeenPurchased.String()))

	request := httptest.NewRequest("PUT", "/api/carritos/3/comprar", &body)
	request.Header.Add("Content-type", "application/json")

	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusPreconditionFailed, response.Code)
	mockPurchaseCarrito.AssertNumberOfCalls(t, "Execute", 1)
}

func createServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	mockPurchaseCarrito = new(mocks.MockPurchaseCarrito)
	handler := PurchaseCarrito{
		UseCase: mockPurchaseCarrito,
	}
	r := gin.Default()
	group := r.Group("/api/carritos")
	group.PUT(":id/comprar", handler.Handle())
	return r
}

func getMockBodySuccess() contract.Body {
	return contract.Body{
		MedioDePagoID: mockMedioDePagoID,
		CompradorID:   mockCompradorID,
	}
}

func getMockCommand() command.PurchaseCarritoCommand {
	return command.PurchaseCarritoCommand{
		CarritoID:     3,
		MedioDePagoID: mockMedioDePagoID,
		CompradorID:   mockCompradorID,
	}
}

func getMockExpectedResponse() []byte {
	resp := response.Response{
		NuevoCarritoID: 3,
	}
	responseJson, _ := json.Marshal(resp)
	return responseJson
}
