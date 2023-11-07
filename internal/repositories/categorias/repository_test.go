package categorias

import (
	"reflect"
	"testing"

	dbmocks "Mileyman-API/internal/app/config/database/mocks"
	"Mileyman-API/internal/domain/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	repository Repository
	mockDB     sqlmock.Sqlmock
	DB         *gorm.DB
)

const (
	QuerySelectByCode = "Call GetCategoriasByDulceID(?)"
	QuerySelectAll    = "SELECT * FROM `Categorias`"
)

func TestGetBycodeOK(t *testing.T) {
	initialize()

	mockCategorias := getMockCategorias()

	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(mockCategorias[0].ID, mockCategorias[0].Nombre).
			AddRow(mockCategorias[1].ID, mockCategorias[1].Nombre))
	response, err := repository.GetCategoriasByDulceID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockCategorias, response)
}

func TestByCodeInternalServerError(t *testing.T) {
	initialize()

	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(1).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetCategoriasByDulceID(1)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func TestGetAllOK(t *testing.T) {
	initialize()

	categorias := getMockCategorias()

	mockDB.ExpectQuery(QuerySelectAll).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(categorias[0].ID, categorias[0].Nombre).
			AddRow(categorias[1].ID, categorias[1].Nombre),
	)
	categoriasRecibidas, err := repository.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, categorias, categoriasRecibidas)
}

func TestGetAllInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectAll).WillReturnError(gorm.ErrInvalidData)

	categoriasRecibidas, err := repository.GetAll()

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, categoriasRecibidas)
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getMockCategorias() (categorias []entities.Categoria) {
	categorias = []entities.Categoria{
		{
			ID:     1,
			Nombre: "Gomitas",
		},
		{
			ID:     2,
			Nombre: "Chocolates",
		},
	}
	return
}
