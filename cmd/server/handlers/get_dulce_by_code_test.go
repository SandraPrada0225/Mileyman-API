package handlers

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"Mileyman-API/internal/use_case/get_dulce_by_code/mocks"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	mockGetDulceByCode *mocks.MockGetDulceByCode
)

func CreateServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	mockGetDulceByCode = new(mocks.MockGetDulceByCode)
	handler := GetDulceByCode{
		UseCase: mockGetDulceByCode,
	}
	r := gin.Default()
	group := r.Group("/api/dulces")
	group.GET("/:codigo", handler.Handle())
	return r
}

func TestOk(t *testing.T) {
	r := CreateServer()
	dulce := entities.Dulce{
		ID:             2,
		Nombre:         "Chocolatina",
		PresentacionID: 1,
		Descripcion:    "Deliciosa chocolatina que se derrite en tu boca",
		Imagen:         "imagen",
		Disponibles:    100,
		Precio:         1000,
		Peso:           40,
		MarcaID:        1,
		Codigo:         "2Mile",
	}

	dulcejson, _ := json.Marshal(&dulce)
	mockGetDulceByCode.On("Execute", "2Mile").Return(dulce, nil)
	request := httptest.NewRequest("GET", "/api/dulces/2Mile", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()
	
	r.ServeHTTP(response, request)

	bodyResponse, err := io.ReadAll(response.Body) 

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(dulcejson), string(bodyResponse))
	mockGetDulceByCode.AssertNumberOfCalls(t, "Execute", 1)

}
func TestWhenNotFoundShouldReturn404(t *testing.T) {
	r := CreateServer()

	mockGetDulceByCode.On("Execute", "2Mile").Return(entities.Dulce{}, database.NewNotFoundError(""))
	request := httptest.NewRequest("GET", "/api/dulces/2Mile", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()
	
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
	mockGetDulceByCode.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenInternalServerErrorShouldReturn500(t *testing.T) {
	r := CreateServer()

	mockGetDulceByCode.On("Execute", "2Mile").Return(entities.Dulce{}, database.NewInternalServerError(""))
	request := httptest.NewRequest("GET", "/api/dulces/2Mile", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()
	
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockGetDulceByCode.AssertNumberOfCalls(t, "Execute", 1)
}

