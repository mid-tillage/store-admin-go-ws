package routers

import (
	"example.com/store-admin-go-ws/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductOnSaleRouter(router *gin.RouterGroup, db *gorm.DB) {
	productOnSaleController := controllers.NewProductOnSaleController(db)

	productOnSaleGroup := router.Group("/product-on-sale")
	{
		productOnSaleGroup.GET("", productOnSaleController.GetProductOnSales)
		productOnSaleGroup.GET("/:id", productOnSaleController.GetProductOnSaleByID)
		productOnSaleGroup.POST("", productOnSaleController.PostProductOnSale)
		productOnSaleGroup.PUT("/:id", productOnSaleController.PutProductOnSale)
		productOnSaleGroup.DELETE("/:id", productOnSaleController.DeleteProductOnSale)
	}
}
