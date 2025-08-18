package main

import (
	"github.com/gin-gonic/gin"
	"goapi/controller"
)

func main() {

	server := gin.Default()

	ProductController := controller.NewProductController()
	//server.GET("/products", ProductController.GetProducts)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}