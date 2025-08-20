package controller

import (
    "github.com/gin-gonic/gin"
    "goapi/usecase"
    "net/http"
)

    type productController struct {
        productUseCase usecase.ProductUseCase
}

    func NewProductController(useCase usecase.ProductUseCase) productController {
        return productController{
            productUseCase: useCase,
        }
    }


func (p *productController) GetProducts(ctx *gin.Context) {
    products, err := p.productUseCase.GetProducts()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to retrieve products",
        })
        return
    }

    ctx.JSON(http.StatusOK, products)
}