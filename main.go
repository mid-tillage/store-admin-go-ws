package main

import (
	docs "example.com/store-admin-go-ws/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"example.com/store-admin-go-ws/routers"
)

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

	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	routers.SetupCatalogRouter(&router.RouterGroup)
	routers.SetupEnterpriseRouter(&router.RouterGroup)
	routers.SetupProductRouter(&router.RouterGroup)
	routers.SetupProductOnSaleRouter(&router.RouterGroup)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3040")
}
