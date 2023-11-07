package getdulcebycode


import (
	"net/http"

	"Mileyman-API/internal/domain/dto/query"
	"Mileyman-API/internal/domain/errors/database"

	"github.com/gin-gonic/gin"
)

type GetDulceByCode struct {
	UseCase UseCase
}

type UseCase interface {
	Execute(codigo string) (query.DetalleDulce, error)
}

func (handler GetDulceByCode) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		codigo := c.Param("codigo")
		dulce, err := handler.UseCase.Execute(codigo)
		if err != nil {
			switch err.(type) {
			case database.NotFoundError:
				c.JSON(http.StatusNotFound, err.Error())
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
			}
			return
		}
		c.JSON(http.StatusOK, dulce)
	}
}
