package updatecarrito

import (
	updatecarrito "Mileyman-API/internal/domain/dto/command/update_carrito"
	"Mileyman-API/internal/domain/dto/query"
	errormessages "Mileyman-API/internal/domain/errors/error_messages"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCarrito struct {
	UseCase UpdateCarritoUseCase
}

type UpdateCarritoUseCase interface {
	Execute(carritoID uint64, movements updatecarrito.Body) query.MovementsResult
}

func (handler UpdateCarrito) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err :=strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, errormessages.IdMustBeAPositiveNumber.String())
			return
		}

		var command updatecarrito.Body
		
		err = c.ShouldBindJSON(&command)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		query := handler.UseCase.Execute(id, command);

		c.JSON(http.StatusOK, query)
	}

}
