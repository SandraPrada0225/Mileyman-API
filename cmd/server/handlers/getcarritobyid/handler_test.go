package getcarritobyid

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
	"Mileyman-API/internal/usecase/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockGetCarritoByID *mocks.MockGetCarritoByID

const (
	mockCarritoID = uint64(3213)
	mockDulceID1  = uint64(3432)
	mockDulceID2  = uint64(3543232)
)

func CreateServer() *gin.Engine {
	gin.SetMode(gin.TestMode)

	mockGetCarritoByID = new(mocks.MockGetCarritoByID)
	handler := GetDetalleCarritoById{
		UseCase: mockGetCarritoByID,
	}
	r := gin.Default()
	group := r.Group("/api/carritos")
	group.GET("/:id", handler.Handle())
	return r
}

func TestOk(t *testing.T) {
	r := CreateServer()
	expectedResponse := responses.GetDetalleCarrito{
		ID:          mockCarritoID,
		Subtotal:    5,
		PrecioTotal: 100,
		Descuento:   5,
		Envio:       5,
		DulcesList: []responses.DulceInCarrito{
			{
				DetalleDulce: responses.DetalleDulce{
					ID:           mockDulceID1,
					Nombre:       "Chocolatina",
					Descripcion:  "Deliciosa chocolatina que se derrite en tu boca",
					Imagen:       "imagen",
					Disponibles:  100,
					PrecioUnidad: 1000,
					Peso:         40,
					Codigo:       "2",
					Presentacion: entities.Presentacion{
						ID:     1,
						Nombre: "Empaque",
					},
					Marca: entities.Marca{
						ID:     2,
						Nombre: "Jet",
					},
					Categorias: []entities.Categoria{
						{
							ID:     2,
							Nombre: "Chocolates",
						},
					},
				},
				Unidades: 2,
				Subtotal: 2000,
			},
			{
				DetalleDulce: responses.DetalleDulce{
					ID:           mockDulceID2,
					Nombre:       "Gomitas",
					Descripcion:  "Ositos de gomita con sabores explosivos",
					Imagen:       "imagen",
					Disponibles:  100,
					PrecioUnidad: 1000,
					Peso:         40,
					Codigo:       "2",
					Presentacion: entities.Presentacion{
						ID:     1,
						Nombre: "Paquete",
					},
					Marca: entities.Marca{
						ID:     2,
						Nombre: "Trululú",
					},
					Categorias: []entities.Categoria{
						{
							ID:     1,
							Nombre: "Gomitas",
						},
					},
				},
				Unidades: 1,
				Subtotal: 1000,
			},
		},
	}

	json, _ := json.Marshal(&expectedResponse)

	mockGetCarritoByID.On("Execute", mockCarritoID).Return(expectedResponse, nil)
	request := httptest.NewRequest("GET", "/api/carritos/3213", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	bodyResponse, err := io.ReadAll(response.Body)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(json), string(bodyResponse))
	mockGetCarritoByID.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenNotFoundShouldReturn404(t *testing.T) {
	r := CreateServer()

	mockGetCarritoByID.On("Execute", mockCarritoID).Return(entities.Dulce{}, database.NewNotFoundError(""))
	request := httptest.NewRequest("GET", "/api/carritos/3213", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
	mockGetCarritoByID.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenInternalServerErrorShouldReturn500(t *testing.T) {
	r := CreateServer()

	mockGetCarritoByID.On("Execute", mockCarritoID).Return(entities.Dulce{}, database.NewInternalServerError(""))
	request := httptest.NewRequest("GET", "/api/carritos/3213", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusInternalServerError, response.Code)
	mockGetCarritoByID.AssertNumberOfCalls(t, "Execute", 1)
}

func TestWhenIDIsInvalidShouldReturn404(t *testing.T) {
	r := CreateServer()

	request := httptest.NewRequest("GET", "/api/carritos/3213a", bytes.NewBuffer([]byte("")))
	request.Header.Add("Content-type", "application/json")
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	mockGetCarritoByID.AssertNumberOfCalls(t, "Execute", 0)
}
