package controller

import (
	"goapi/model"
	"goapi/usecase"
	"net/http"
    "strconv"
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

func (p *productController) GetProductByID(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID, need to be a number"})
        return
    }
    product, err := p.productUseCase.GetProductByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    ctx.JSON(http.StatusOK, product)
}

func (p *productController) GetProductByName(ctx *gin.Context) {
    name := ctx.Query("name")

    if name == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name query parameter is required"})
    }

    products, err := p.productUseCase.GetProductByName(name)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching product"})
        return
    }
    ctx.JSON(http.StatusOK, products)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
    idParam := ctx.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    var product model.Product
    if err := ctx.BindJSON(&product); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }
    product.ID = id
    if err := p.productUseCase.UpdateProduct(product); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product"})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}