package usuarios

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
	UpdateQuery = "UPDATE `usuarios` SET `nombre`=?,`apellido`=?,`password`=?,`correo`=?,`carrito_actual_id`=? WHERE `id` = ?"
	SelectQuery = "SELECT * FROM `usuarios` WHERE id = ? LIMIT 1"
)

func TestSaveWhenIsSuccessfullShouldReturnNoError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(UpdateQuery).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()
	mockUsuario := getMockUsuario()

	err := repository.Save(&mockUsuario)

	assert.NoError(t, err)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestSaveWhenWentWrongShouldReturnInternalError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(UpdateQuery).WillReturnError(gorm.ErrInvalidData)
	mockDB.ExpectRollback()
	mockUsuario := getMockUsuario()

	err := repository.Save(&mockUsuario)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
}

func TestSaveWhenCarBelongsToOtherUserShouldReturnConflictError(t *testing.T) {
	initialize()
	mockDB.ExpectBegin()
	mockDB.ExpectExec(UpdateQuery).WillReturnError(gorm.ErrDuplicatedKey)
	mockDB.ExpectRollback()
	mockUsuario := getMockUsuario()

	err := repository.Save(&mockUsuario)

	assert.Error(t, err)
	assert.Equal(t, "database.ConflictError", reflect.TypeOf(err).String())
}

func TestGetByIDWhenIsSuccessfullReturnUsuario(t *testing.T) {
	initialize()
	mockUsuario := getMockUsuario()

	mockDB.ExpectQuery(SelectQuery).WithArgs(mockUsuario.ID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "nombre", "apellido", "correo", "password", "carrito_actual_id"}).
			AddRow(mockUsuario.ID, mockUsuario.Nombre, mockUsuario.Apellido, mockUsuario.Correo, mockUsuario.Password, mockUsuario.CarritoActualID),
	)

	usuario, err := repository.GetByID(mockUsuario.ID)

	assert.NoError(t, err)
	assert.Equal(t, mockUsuario, usuario)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetByIDWhenUsuarioDoesNotExistsShouldReturnNotFoundError(t *testing.T) {
	initialize()
	mockUsuario := getMockUsuario()

	mockDB.ExpectQuery(SelectQuery).WithArgs(mockUsuario.ID).WillReturnError(gorm.ErrRecordNotFound)

	usuario, err := repository.GetByID(mockUsuario.ID)

	assert.Error(t, err)
	assert.Equal(t, "database.NotFoundError", reflect.TypeOf(err).String())
	assert.Empty(t, usuario)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestGetByIDWhenWentWrongShouldReturnInternalError(t *testing.T) {
	initialize()
	mockUsuario := getMockUsuario()

	mockDB.ExpectQuery(SelectQuery).WithArgs(mockUsuario.ID).WillReturnError(gorm.ErrInvalidData)

	usuario, err := repository.GetByID(mockUsuario.ID)

	assert.Error(t, err)
	assert.Equal(t, "database.InternalServerError", reflect.TypeOf(err).String())
	assert.Empty(t, usuario)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func getMockUsuario() entities.Usuario {
	return entities.Usuario{
		ID:              1,
		Nombre:          "Frey",
		Apellido:        "Man",
		CarritoActualID: 2,
		Password:        "MeQuieroMuch0",
	}
}

func initialize() {
	mockDB, DB = dbmocks.NewDB()
	mockDB.MatchExpectationsInOrder(false)
	repository = Repository{
		DB: DB.Debug(),
	}
}
