package main

import (
	"github.com/gin-gonic/gin"
	"goapi/controller"
	"goapi/db"
	"goapi/repository"
	"goapi/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada de repositories
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada de usecases
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)
	//server.GET("/products", ProductController.GetProducts)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}