package main

import (
	"encoding/json"
	"fmt"

	docs "example.com/store-admin-go-ws/docs"
	"example.com/store-admin-go-ws/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"example.com/store-admin-go-ws/routers"
)

var db *gorm.DB // Global variable for the database connection

// @title           store-admin-go-ws API
// @version         1.0
// @description     This is a store admin WS sample server.
// @termsOfService  -

// @contact.name   Mid-Tillage
// @contact.url    https://github.com/mid-tillage/store-admin-go-ws
// @contact.email  ripioconarcilla@gmail.com

// @license.name  -
// @license.url   -

// @host      localhost:3040
// @BasePath

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=sale-management-system port=5432 sslmode=disable TimeZone=America/Santiago"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	routers.SetupCatalogRouter(&router.RouterGroup)
	routers.SetupEnterpriseRouter(&router.RouterGroup)
	routers.SetupProductRouter(&router.RouterGroup)
	routers.SetupProductOnSaleRouter(&router.RouterGroup, db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// var catalogs []models.Catalog
	// db.Table("commerce.dim_product_catalog").Find(&catalogs)
	// catalogsJSON, err := json.Marshal(catalogs)
	// if err != nil {
	// 	panic("failed to marshal catalogs to JSON")
	// }
	// fmt.Println("Catalogs:" + string(catalogsJSON))

	type Test struct {
		Id_product    uint
		Name          string
		Id_enterprise uint
		Ename         string
	}
	// var product []models.Product
	// db.Table("commerce.dim_product").
	// 	Select("dim_product.id_product as id_product, dim_product.name as name, dim_product.id_enterprise as id_enterprise, dim_enterprise.name as ename").
	// 	Joins("inner join commerce.dim_enterprise on commerce.dim_product.id_enterprise = commerce.dim_enterprise.id_enterprise").
	// 	Where("dim_product.id_product = 1").
	// 	Scan(&product)
	// db.Table("commerce.dim_product").
	// 	Select("id_product, dim_product.name, dim_product.id_enterprise, dim_enterprise.name as ename").
	// 	Joins("inner join commerce.dim_enterprise on dim_product.id_enterprise = dim_enterprise.id_enterprise").
	// 	Where("id_product = 1").Scan(&product)
	// db.Joins("Enterprise").Where(&models.Product{IDProduct: 1}).Find(&product)
	// productJson, err := json.Marshal(product)
	// fmt.Println("Products:" + string(productJson))

	var pos []models.ProductOnSale
	// db.Joins("Product").Where(&models.ProductOnSale{IdProductOnSale: 16}).Find(&pos)
	db.Table("commerce.fact_product_on_sale").
		Joins("inner join commerce.dim_product on fact_product_on_sale.id_product = dim_product.id_product").
		Where("fact_product_on_sale.id_product_on_sale = 16").
		Find(&pos)
	posJson, err := json.Marshal(pos)
	fmt.Println("POS:" + string(posJson))
	// var enterprises []models.Enterprise
	// db.Table("commerce.dim_enterprise").Find(&enterprises)
	// enterprisesJSON, err := json.Marshal(enterprises)
	// if err != nil {
	// 	panic("failed to marshal enterprises to JSON")
	// }
	// fmt.Println("Enterprises:" + string(enterprisesJSON))

	// var products []models.Product
	// // db.Table("commerce.dim_product").Find(&products)
	// db.Preload("Product").Preload("Enterprise").Find(&products)
	// productsJSON, err := json.Marshal(products)
	// if err != nil {
	// 	panic("failed to marshal products to JSON")
	// }
	// fmt.Println("Products:" + string(productsJSON))

	router.Run(":3040")
}
