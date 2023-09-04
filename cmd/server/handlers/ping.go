package handlers

import "github.com/gin-gonic/gin"

type Ping struct{}

func (handler Ping) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "Pong")
	}
}
