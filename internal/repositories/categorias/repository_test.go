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
)

func TestGetBycodeOK(t *testing.T) {
	inicialize()

	mockCategorias := GetMockCategorias()

	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(1).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(mockCategorias[0].ID, mockCategorias[0].Nombre).
			AddRow(mockCategorias[1].ID, mockCategorias[1].Nombre))
	response, err := repository.GetCategoriasByDulceID(1)

	assert.NoError(t, err)
	assert.Equal(t, mockCategorias, response)
}

func TestByCodeInternalServerError(t *testing.T) {
	inicialize()
	mockDB.ExpectQuery(QuerySelectByCode).WithArgs(1).WillReturnError(gorm.ErrInvalidData)

	dulceRecibido, err := repository.GetCategoriasByDulceID(1)

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, dulceRecibido)
}

func inicialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func GetMockCategorias() (categorias []entities.Categoria) {
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
