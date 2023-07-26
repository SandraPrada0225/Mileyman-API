package main

import (
	"Mileyman-API/cmd/server/handlers"
	"Mileyman-API/internal/app/config/database"
	"Mileyman-API/internal/domain/entities"
	"Mileyman-API/internal/repositories/dulces"
	"fmt"

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
	repositorio := dulces.Repository{
		DB: db,
	}
	var dulce entities.Dulce

	dulce, err = repositorio.GetByCode("12")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n",dulce)

	ping_handler := handlers.Ping{}

	// Ruta de la API
	r.GET("/ping", ping_handler.Handle())

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}

func handlePing(c *gin.Context) {
	c.String(200, "pong")
}
