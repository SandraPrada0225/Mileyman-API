package getfiltros

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	mocks "Mileyman-API/internal/usecase/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockGetFiltros *mocks.MockGetFiltros

func CreateServerGetFiltros() *gin.Engine {
	gin.SetMode(gin.TestMode)

	mockGetFiltros = new(mocks.MockGetFiltros)
	handler := GetFiltros{
		UseCase: mockGetFiltros,
	}
	r := gin.Default()
	group := r.Group("/api/filtros")
	group.GET("/", handler.Handle())
	return r
}

func TestOkGetFiltros(t *testing.T) {
	r := CreateServerGetFiltros()

	filtros := responses.GetFiltros{
		Categorias: []entities.Categoria{
			{
				ID:     1,
				Nombre: "Gomitas",
			},
			{
				ID:     2,
				Nombre: "Chupetes",
			},
		},
		Marcas: []entities.Marca{
			{
				ID:     1,
				Nombre: "Trululu",
			},
			{
				ID:     2,
				Nombre: "Jet",
			},
		},
		Presentaciones: []entities.Presentacion{
			{
				ID:     1,
				Nombre: "Caja",
			},
			{
				ID:     2,
				Nombre: "Bolsa",
			},
		},
	}

	filtrosjson, _ := json.Marshal(&filtros)
	mockGetFiltros.On("Execute").Return(filtros, nil)
	request := httptest.NewRequest("GET", "/api/filtros/", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	bodyResponse, err := io.ReadAll(response.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(filtrosjson), string(bodyResponse))
	mockGetFiltros.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenInternalServerErrorShouldReturn500GetFiltros(t *testing.T) {
	r := CreateServerGetFiltros()

	mockGetFiltros.On("Execute").Return(responses.GetFiltros{}, database.NewInternalServerError(""))
	request := httptest.NewRequest("GET", "/api/filtros/", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockGetFiltros.AssertNumberOfCalls(t, "Execute", 1)
}
