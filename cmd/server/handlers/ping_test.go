package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CreateServerPing() *gin.Engine {
	gin.SetMode(gin.TestMode)

	handler := Ping{}
	r := gin.Default()
	r.GET("/ping", handler.Handle())
	return r
}

func TestPing(t *testing.T) {
	r := CreateServerPing()
	request := httptest.NewRequest("GET", "/ping", bytes.NewBuffer([]byte("")))

	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}
