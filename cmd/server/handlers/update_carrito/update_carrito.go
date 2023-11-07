package updatecarrito

import (
	"net/http"
	"strconv"

	updatecarrito "Mileyman-API/internal/domain/dto/command/update_carrito"
	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"

	"github.com/gin-gonic/gin"
)

type UpdateCarrito struct {
	UseCase UpdateCarritoUseCase
}

type UpdateCarritoUseCase interface {
	Execute(carritoID uint64, movements updatecarrito.Body) (query.MovementsResult, error)
}

func (handler UpdateCarrito) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errormessages.IdMustBeAPositiveNumber.String())
			return
		}

		var command updatecarrito.Body

		err = c.ShouldBindJSON(&command)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		query, err := handler.UseCase.Execute(id, command)
		if err != nil {
			switch err.(type) {
			case database.NotFoundError:
				c.JSON(http.StatusNotFound, err.Error())
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}

		c.JSON(http.StatusOK, query)
	}
}
