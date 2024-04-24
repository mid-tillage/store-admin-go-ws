```go
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Catalog struct {
	IDProductCatalog int64
	Name             string
}

type Enterprise struct {
	IDEnterprise int64
	Name         string
}

type Product struct {
	Name        string
	Description string
	Enterprise  Enterprise
}

type ProductOnSale struct {
	IDProductOnSale   int64
	Title             string
	Product           Product
	Catalog           Catalog
	Price             float64
	SaleStartDatetime time.Time
	SaleEndDatetime   time.Time
}

var productsOnSale = []ProductOnSale{
	{
		IDProductOnSale: 1,
		Title:           "Product 1",
		Product: Product{
			Name:        "Product 1",
			Description: "Description 1",
			Enterprise: Enterprise{
				IDEnterprise: 1,
				Name:         "Enterprise 1",
			},
		},
		Catalog: Catalog{
			IDProductCatalog: 1,
			Name:             "Catalog 1",
		},
		Price:             10.5,
		SaleStartDatetime: time.Now(),
		SaleEndDatetime:   time.Now().AddDate(0, 1, 0),
	},
	{
		IDProductOnSale: 2,
		Title:           "Product 2",
		Product: Product{
			Name:        "Product 2",
			Description: "Description 2",
			Enterprise: Enterprise{
				IDEnterprise: 2,
				Name:         "Enterprise 2",
			},
		},
		Catalog: Catalog{
			IDProductCatalog: 2,
			Name:             "Catalog 2",
		},
		Price:             20.75,
		SaleStartDatetime: time.Now(),
		SaleEndDatetime:   time.Now().AddDate(0, 2, 0),
	},
}

func main() {
	router := gin.Default()

	router.GET("/product-on-sale", getProductOnSales)
	router.GET("/product-on-sale/:id", getProductOnSaleByID)
	router.POST("/product-on-sale", postProductOnSale)
	router.PUT("/product-on-sale/:id", putProductOnSale)
	router.DELETE("/product-on-sale/:id", deleteProductOnSale)

	router.Run(":3040")
}

func getProductOnSales(c *gin.Context) {
	c.JSON(http.StatusOK, productsOnSale)
}

func getProductOnSaleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var foundProductOnSale *ProductOnSale
	for _, a := range productsOnSale {
		if a.IDProductOnSale == id {
			foundProductOnSale = &a
			break
		}
	}
	if foundProductOnSale == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, foundProductOnSale)
}

func postProductOnSale(c *gin.Context) {
	var newProductOnSale ProductOnSale
	fmt.Println("Received new product on sale:", newProductOnSale)

	if err := c.BindJSON(&newProductOnSale); err != nil {
		fmt.Println("Error binding JSON:", err)
		return
	}

	productsOnSale = append(productsOnSale, newProductOnSale)
	c.IndentedJSON(http.StatusCreated, newProductOnSale)
}

func putProductOnSale(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedProductOnSale ProductOnSale
	if err := c.BindJSON(&updatedProductOnSale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	var index int = -1
	for i, p := range productsOnSale {
		if p.IDProductOnSale == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	productsOnSale[index] = updatedProductOnSale
	c.JSON(http.StatusOK, updatedProductOnSale)
}

func deleteProductOnSale(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var index int = -1
	for i, p := range productsOnSale {
		if p.IDProductOnSale == id {
			index = i
			break
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	productsOnSale = append(productsOnSale[:index], productsOnSale[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
```