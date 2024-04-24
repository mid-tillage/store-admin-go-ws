// controllers/product_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/services"
)

type ProductOnSaleController struct {
	productOnSaleService *services.ProductOnSaleService
}

func NewProductOnSaleController() *ProductOnSaleController {
	return &ProductOnSaleController{
		productOnSaleService: services.NewProductOnSaleService(),
	}
}

func (pc *ProductOnSaleController) GetProductOnSales(c *gin.Context) {
	// Implementation...
}

func (pc *ProductOnSaleController) GetProductOnSaleByID(c *gin.Context) {
	// Implementation...
}

func (pc *ProductOnSaleController) PostProductOnSale(c *gin.Context) {
	// Implementation...
}

func (pc *ProductOnSaleController) PutProductOnSale(c *gin.Context) {
	// Implementation...
}

func (pc *ProductOnSaleController) DeleteProductOnSale(c *gin.Context) {
	// Implementation...
}
