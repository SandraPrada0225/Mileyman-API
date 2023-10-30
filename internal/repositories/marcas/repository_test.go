package marcas

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
	QuerySelectAll = "SELECT * FROM `Marcas`"
)

func TestGetAllOK(t *testing.T) {
	initialize()

	marcas := getMarcas()

	mockDB.ExpectQuery(QuerySelectAll).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre"}).
			AddRow(marcas[0].ID, marcas[0].Nombre).
			AddRow(marcas[1].ID, marcas[1].Nombre),
	)
	marcasRecibidas, err := repository.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, marcas, marcasRecibidas)
}

func TestByCodeInternalServerError(t *testing.T) {
	initialize()
	mockDB.ExpectQuery(QuerySelectAll).WillReturnError(gorm.ErrInvalidData)

	marcasRecibidas, err := repository.GetAll()

	typeErr := reflect.TypeOf(err).String()

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", typeErr)
	assert.Empty(t, marcasRecibidas)
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}

func getMarcas() (marcas []entities.Marca) {
	marcas = []entities.Marca{
		{
			ID:     1,
			Nombre: "Trululu",
		},
		{
			ID:     2,
			Nombre: "Jet",
		},
	}

	return
}
