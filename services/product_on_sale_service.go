// services/product_service.go
package services

import (
	"errors"
	"time"

	"example.com/store-admin-go-ws/models"
)

type ProductOnSaleService struct {
	// Add any dependencies here...
}

func NewProductOnSaleService() *ProductOnSaleService {
	return &ProductOnSaleService{}
}

var productsOnSale = []models.ProductOnSale{
	{
		IDProductOnSale: 1,
		Title:           "Product 1",
		Product: models.Product{
			Name:        "Product 1",
			Description: "Description 1",
			Enterprise: models.Enterprise{
				IDEnterprise: 1,
				Name:         "Enterprise 1",
			},
		},
		Catalog: models.Catalog{
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
		Product: models.Product{
			Name:        "Product 2",
			Description: "Description 2",
			Enterprise: models.Enterprise{
				IDEnterprise: 2,
				Name:         "Enterprise 2",
			},
		},
		Catalog: models.Catalog{
			IDProductCatalog: 2,
			Name:             "Catalog 2",
		},
		Price:             20.75,
		SaleStartDatetime: time.Now(),
		SaleEndDatetime:   time.Now().AddDate(0, 2, 0),
	},
}

func (ps *ProductOnSaleService) GetProductOnSales() ([]models.ProductOnSale, error) {
	return productsOnSale, nil
}

func (ps *ProductOnSaleService) GetProductOnSaleByID(id int64) (*models.ProductOnSale, error) {
	for _, p := range productsOnSale {
		if p.IDProductOnSale == id {
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}

func (ps *ProductOnSaleService) PostProductOnSale(productOnSale *models.ProductOnSale) error {
	// Assign a new ID to the product
	productOnSale.IDProductOnSale = generateNextID()

	// Add the new product to the list
	productsOnSale = append(productsOnSale, *productOnSale)
	return nil
}

func (ps *ProductOnSaleService) PutProductOnSale(id int64, updatedProductOnSale *models.ProductOnSale) error {
	// Find the index of the product with the given ID
	index := -1
	for i, p := range productsOnSale {
		if p.IDProductOnSale == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("product not found")
	}

	// Update the product
	productsOnSale[index] = *updatedProductOnSale
	return nil
}

func (ps *ProductOnSaleService) DeleteProductOnSale(id int64) error {
	index := -1
	for i, p := range productsOnSale {
		if p.IDProductOnSale == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("product not found")
	}

	// Remove the product from the list
	productsOnSale = append(productsOnSale[:index], productsOnSale[index+1:]...)
	return nil
}

func generateNextID() int64 {
	// Find the maximum ID currently in use
	maxID := int64(0)
	for _, p := range productsOnSale {
		if p.IDProductOnSale > maxID {
			maxID = p.IDProductOnSale
		}
	}
	// Increment the maximum ID by 1 to get the next available ID
	return maxID + 1
}
