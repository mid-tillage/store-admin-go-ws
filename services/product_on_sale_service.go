// services/product_service.go
package services

import (
	"net/http"
	"time"

	"github.com/mid-tillage/store-admin-go-ws/models"
)

type ProductService struct {
	// Add any dependencies here...
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

func NewProductService() *ProductService {
	return &ProductService{}
}

func (ps *ProductService) GetProductOnSales() ([]models.ProductOnSale, error) {
	c.JSON(http.StatusOK, productsOnSale)
}

func (ps *ProductService) GetProductOnSaleByID(id int64) (*models.ProductOnSale, error) {
	// Implementation...
}

func (ps *ProductService) PostProductOnSale(productOnSale *models.ProductOnSale) error {
	// Implementation...
}

func (ps *ProductService) PutProductOnSale(id int64, updatedProductOnSale *models.ProductOnSale) error {
	// Implementation...
}

func (ps *ProductService) DeleteProductOnSale(id int64) error {
	// Implementation...
}
