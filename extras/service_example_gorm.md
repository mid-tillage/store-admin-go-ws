```go
package services

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"example.com/store-admin-go-ws/models"
)

type ProductOnSaleService struct {
	db *gorm.DB
}

func NewProductOnSaleService(db *gorm.DB) *ProductOnSaleService {
	return &ProductOnSaleService{db: db}
}

func (ps *ProductOnSaleService) GetProductOnSales() ([]models.ProductOnSale, error) {
	var productsOnSale []models.ProductOnSale
	if err := ps.db.
		Preload("Product").
		Preload("Catalog").
		Preload("Product.Enterprise").
		Table("commerce.fact_product_on_sale").
		Joins("inner join commerce.dim_product on fact_product_on_sale.id_product = dim_product.id_product").
		Joins("inner join commerce.dim_enterprise on dim_product.id_enterprise = dim_enterprise.id_enterprise").
		Joins("inner join commerce.dim_product_catalog on fact_product_on_sale.id_product_catalog = dim_product_catalog.id_product_catalog").
		Find(&productsOnSale).Error; err != nil {
		fmt.Println("Error fetching products:", err)
		return nil, err
	}
	return productsOnSale, nil
}

func (ps *ProductOnSaleService) GetProductOnSaleByID(id int64) (*models.ProductOnSale, error) {
	var productOnSale models.ProductOnSale
	if err := ps.db.Model(&models.ProductOnSale{}).First(&productOnSale, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &productOnSale, nil
}

func (ps *ProductOnSaleService) PostProductOnSale(productOnSale *models.ProductOnSale) error {
	return ps.db.Omit("Product").Omit("Catalog").Create(productOnSale).Error
}

func (ps *ProductOnSaleService) PutProductOnSale(id int64, updatedProductOnSale *models.ProductOnSale) error {
	if err := ps.db.Model(&models.ProductOnSale{}).Where(&models.ProductOnSale{IdProductOnSale: id}).Updates(updatedProductOnSale).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}
	return nil
}

func (ps *ProductOnSaleService) DeleteProductOnSale(id int64) error {
	if err := ps.db.Delete(&models.ProductOnSale{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}
	return nil
}

```