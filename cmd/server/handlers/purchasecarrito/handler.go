package purchasecarrito

import (
	"net/http"
	"strconv"

	command "Mileyman-API/internal/domain/dto/command/purchasecarrito"
	contract "Mileyman-API/internal/domain/dto/contracts/purchasecarrito"
	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/errors"
	errormessages "Mileyman-API/internal/domain/errors/errormessages"
	"Mileyman-API/internal/domain/errors/rest"

	"github.com/gin-gonic/gin"
)

type GetFiltrosUseCase interface {
	Execute(command command.PurchaseCarritoCommand) (uint64, error)
}

type PurchaseCarrito struct {
	UseCase GetFiltrosUseCase
}

func (handler PurchaseCarrito) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request contract.Request

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			params := errormessages.Parameters{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, rest.NewBadRequestError(errormessages.InvalidTypeError.GetMessageWithParams(params)).Error())
			return
		}

		request.URLParams.CarritoID = id

		err = c.ShouldBindJSON(&request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		command := command.NewPurchaseCarritoCommandFromRequest(request)

		newCarritoID, err := handler.UseCase.Execute(command)
		if err != nil {
			errors.GetAPIErrors(err, c)
			return
		}

		c.JSON(http.StatusOK, responses.NewResponse(newCarritoID))
	}
}
