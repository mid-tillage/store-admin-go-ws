package routers

import (
	"github.com/gin-gonic/gin"

	"example.com/store-admin-go-ws/controllers"
)

func SetupProductOnSaleRouter(router *gin.RouterGroup) {
	productOnSaleController := controllers.NewProductOnSaleController()

	productOnSaleGroup := router.Group("/product-on-sale")
	{
		productOnSaleGroup.GET("", productOnSaleController.GetProductOnSales)
		productOnSaleGroup.GET("/:id", productOnSaleController.GetProductOnSaleByID)
		productOnSaleGroup.POST("", productOnSaleController.PostProductOnSale)
		productOnSaleGroup.PUT("/:id", productOnSaleController.PutProductOnSale)
		productOnSaleGroup.DELETE("/:id", productOnSaleController.DeleteProductOnSale)
	}
}
