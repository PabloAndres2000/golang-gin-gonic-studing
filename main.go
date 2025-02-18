package main

import (
	"golang-gin-gonic-studing/routes"

	_ "golang-gin-gonic-studing/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Usuarios con Gin
// @version 1.0
// @description En este apartado se almacena APIS de Usuarios utilizando Gin y Swagger
// @host localhost:8080
// @BasePath /
func main() {
	r := routes.SetupRouter()

	// Ruta para la documentación de Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
