package getpurchaselistbyuserid

import (
	"net/http"
	"strconv"

	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/domain/errors/errormessages"
	"Mileyman-API/internal/domain/errors/rest"

	"github.com/gin-gonic/gin"
)

type GetPurchaseListByUserID struct {
	UseCase GetPurchaseListByUserIDUseCase
}

type GetPurchaseListByUserIDUseCase interface {
	Execute(userID uint64) (responses.GetPurchaseList, error)
}

func (handler GetPurchaseListByUserID) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			params := errormessages.Parameters{
				"error": err.Error(),
			}
			c.JSON(http.StatusBadRequest, rest.NewBadRequestError(errormessages.InvalidTypeError.GetMessageWithParams(params)).Error())
			return
		}

		response, err := handler.UseCase.Execute(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, response)
	}
}
