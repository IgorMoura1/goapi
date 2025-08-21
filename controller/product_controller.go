package controller

import (
	"goapi/model"
	"goapi/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (p *productController) CreateProduct(ctx *gin.Context) {

    var product model.Product
    if err := ctx.BindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "Erro ao inserir dados do produto",
        })
        return
    }

    insertedProduct, err := p.productUseCase.CreateProduct(product)
    if err != nil { 
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": "Erro interno ao inserir produto",
        })
        return
    }    
        
    ctx.JSON(http.StatusCreated, insertedProduct)

}