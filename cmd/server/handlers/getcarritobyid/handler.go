package getcarritobyid

import (
	"net/http"
	"strconv"

	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/errors/database"
	errormessages "Mileyman-API/internal/domain/errors/errormessages"
	"Mileyman-API/internal/domain/errors/rest"

	"github.com/gin-gonic/gin"
)

type GetCarritoUseCase interface {
	Execute(id uint64) (responses.GetDetalleCarrito, error)
}

type GetDetalleCarritoById struct {
	UseCase GetCarritoUseCase
}

func (handler GetDetalleCarritoById) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			params := errormessages.Parameters{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, rest.NewBadRequestError(errormessages.InvalidTypeError.GetMessageWithParams(params)).Error())
			return
		}

		carrito, err := handler.UseCase.Execute(id)
		if err != nil {
			switch err.(type) {
			case database.NotFoundError:
				c.JSON(http.StatusNotFound, err.Error())
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}
		c.JSON(http.StatusOK, carrito)
	}
}
