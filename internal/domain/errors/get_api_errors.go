package errors

import (
	"net/http"

	"Mileyman-API/internal/domain/errors/business"
	"Mileyman-API/internal/domain/errors/database"

	"github.com/gin-gonic/gin"
)

func GetAPIErrors(err error, c *gin.Context) {
	switch err.(type) {
	case database.NotFoundError:
		c.JSON(http.StatusNotFound, err.Error())
	case database.InternalServerError:
		c.JSON(http.StatusInternalServerError, err.Error())
	case business.CarritoAlreadyPurchaseError:
		c.JSON(http.StatusPreconditionFailed, err.Error())
	}
}
