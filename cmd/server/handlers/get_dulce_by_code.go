package handlers
// es un camarero

import (
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/domain/errors/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetDulceByCode struct {
	UseCase UseCase
}

type UseCase interface {
	Execute(codigo string) (entities.Dulce, error)
}

func (handler GetDulceByCode) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		codigo := c.Param("codigo")
		dulce, err := handler.UseCase.Execute(codigo)

		if err != nil {
			switch err.(type) {
			case database.NotFoundError:
				c.JSON(http.StatusNotFound, err.Error())
			case database.InternalServerError:
				c.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}
		c.JSON(http.StatusOK, dulce)

	}
}
