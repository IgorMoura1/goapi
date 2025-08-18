package controller

import (
    "github.com/gin-gonic/gin" // <-- add this import
    "goapi/model"
    "net/http"
)

type productController struct {
    // usecase
}

func NewProductController() productController {
    return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
    products := []model.Product{
        {
            ID:    1,
            Name:  "Batata frita",
            Price: 10.0,
        },
    }

    ctx.JSON(http.StatusOK, products)
}