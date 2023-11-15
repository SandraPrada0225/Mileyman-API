package getfiltros

import (
	"net/http"

	"Mileyman-API/internal/domain/dto/responses"

	"github.com/gin-gonic/gin"
)

type GetFiltros struct {
	UseCase GetFiltrosUseCase
}

type GetFiltrosUseCase interface {
	Execute() (responses.GetFiltros, error)
}

func (handler GetFiltros) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		query, err := handler.UseCase.Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, query)
	}
}
