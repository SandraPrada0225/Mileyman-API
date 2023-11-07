package updatecarrito

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	updatecarrito "Mileyman-API/internal/domain/dto/command/update_carrito"
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/errors/database"
	mocks "Mileyman-API/internal/use_case/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	mockCarritoID  = uint64(1)
	mockCarritoID2 = uint64(6)
)

var mockUpdateCarrito *mocks.MockUpdateCarrito

func CreateServerUpdateCarrito() *gin.Engine {
	gin.SetMode(gin.TestMode)

	mockUpdateCarrito = new(mocks.MockUpdateCarrito)
	handler := UpdateCarrito{
		UseCase: mockUpdateCarrito,
	}
	r := gin.Default()
	group := r.Group("/api/carritos")
	group.PUT("/:id", handler.Handle())
	return r
}

func TestOkUpdateCarrito(t *testing.T) {
	r := CreateServerUpdateCarrito()

	movementsResult := query.MovementsResult{
		Result: []query.MovementResult{
			{
				Movement: 0,
				DulceID:  1,
				Result:   "Updated",
				Error:    "",
			},
			{
				Movement: 1,
				DulceID:  2,
				Result:   "Created",
				Error:    "",
			},
			{
				Movement: 2,
				DulceID:  1,
				Result:   "Deleted",
				Error:    "",
			},
		},
	}

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(getMockMovements())
	if err != nil {
		panic(err.Error())
	}

	movementsResultjson, err := json.Marshal(&movementsResult)
	if err != nil {
		panic(err.Error())
	}
	mockUpdateCarrito.On("Execute", mockCarritoID, getMockMovements()).Return(movementsResult, nil)
	request := httptest.NewRequest("PUT", "/api/carritos/1", &body)
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	bodyResponse, err := io.ReadAll(response.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(movementsResultjson), string(bodyResponse))
	mockUpdateCarrito.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenNotFoundShouldReturn404(t *testing.T) {
	r := CreateServerUpdateCarrito()

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(getMockMovements())
	if err != nil {
		panic(err.Error())
	}
	mockUpdateCarrito.On("Execute", mockCarritoID2, getMockMovements()).Return(query.MovementsResult{}, database.NewNotFoundError(""))
	request := httptest.NewRequest("PUT", "/api/carritos/6", &body)
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
	mockUpdateCarrito.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenInternalServerErrorShouldReturn500(t *testing.T) {
	r := CreateServerUpdateCarrito()

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(getMockMovements())
	if err != nil {
		panic(err.Error())
	}
	mockUpdateCarrito.On("Execute", mockCarritoID, getMockMovements()).Return(query.MovementsResult{}, database.NewInternalServerError(""))
	request := httptest.NewRequest("PUT", "/api/carritos/1", &body)
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockUpdateCarrito.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenInternalServerErrorShouldReturn400(t *testing.T) {
	r := CreateServerUpdateCarrito()

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(getMockMovements())
	if err != nil {
		panic(err.Error())
	}
	request := httptest.NewRequest("PUT", "/api/carritos/s", &body)
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestWhenShouldBindJSONFailedShouldReturn400(t *testing.T) {
	r := CreateServerUpdateCarrito()

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode("StatusUnprocessableEntity")
	if err != nil {
		panic(err.Error())
	}
	request := httptest.NewRequest("PUT", "/api/carritos/1", &body)
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func getMockMovements() (movements updatecarrito.Body) {
	movements = updatecarrito.Body{
		Movements: []updatecarrito.Movement{
			{
				DulceID:  1,
				Unidades: 4,
			},
			{
				DulceID:  2,
				Unidades: 2,
			},
			{
				DulceID:  1,
				Unidades: 0,
			},
		},
	}
	return
}
