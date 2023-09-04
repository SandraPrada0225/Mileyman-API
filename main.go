package main

import (
	"Mileyman-API/cmd/server/routes"
	"Mileyman-API/internal/app/config/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Crear una instancia del enrutador de Gin
	r := gin.Default()

	client := database.Client{}
	db, err := client.Connect()
	if err != nil {
		panic(err)
	}

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}

func handlePing(c *gin.Context) {
	c.String(200, "pong")
}
