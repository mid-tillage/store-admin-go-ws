// controllers/product_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"example.com/store-admin-go-ws/models"
	"example.com/store-admin-go-ws/services"
	"github.com/gin-gonic/gin"
)

type ProductOnSaleController struct {
	productOnSaleService *services.ProductOnSaleService
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func NewProductOnSaleController() *ProductOnSaleController {
	return &ProductOnSaleController{}
}

// @Summary Get all products on sale
// @Description Retrieves a list of all products on sale.
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} models.ProductOnSale
// @Router /product-on-sale [get]
func (pc *ProductOnSaleController) GetProductOnSales(c *gin.Context) {
	products, err := pc.productOnSaleService.GetProductOnSales()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// @Summary Get a product on sale by ID
// @Description Retrieves a product on sale by its ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.ProductOnSale
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product-on-sale/{id} [get]
func (pc *ProductOnSaleController) GetProductOnSaleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := pc.productOnSaleService.GetProductOnSaleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Create a new product on sale
// @Description Creates a new product on sale.
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.ProductOnSale true "Product Object"
// @Success 201 {object} models.ProductOnSale
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product-on-sale [post]
func (pc *ProductOnSaleController) PostProductOnSale(c *gin.Context) {
	var product models.ProductOnSale
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.productOnSaleService.PostProductOnSale(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// @Summary Update a product on sale
// @Description Updates a product on sale by its ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.ProductOnSale true "Updated Product Object"
// @Success 200 {object} models.ProductOnSale
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product-on-sale/{id} [put]
func (pc *ProductOnSaleController) PutProductOnSale(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedProduct models.ProductOnSale
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.productOnSaleService.PutProductOnSale(id, &updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// @Summary Delete a product on sale
// @Description Deletes a product on sale by its ID.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product-on-sale/{id} [delete]
func (pc *ProductOnSaleController) DeleteProductOnSale(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := pc.productOnSaleService.DeleteProductOnSale(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
