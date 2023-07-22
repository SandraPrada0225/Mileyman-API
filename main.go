package main

import (
	"Mileyman-API/cmd/server/handlers"

	"github.com/gin-gonic/gin"
)


func main() {
    // Crear una instancia del enrutador de Gin
    r := gin.Default()

	ping_handler := handlers.Ping{}
	
    // Ruta de la API
	r.GET("/ping",ping_handler.Handle())

    // Iniciar el servidor en el puerto 8080
    r.Run(":8080")
}


func handlePing(c *gin.Context){
	c.String(200, "pong")
}