package models

import "time"

type Catalog struct {
	IdProductCatalog uint   `gorm:"primaryKey" json:"idProductCatalog"`
	Name             string `json:"name"`
}

type Enterprise struct {
	IdEnterprise uint   `gorm:"primaryKey" json:"idEnterprise"`
	Name         string `json:"name"`
}

type Product struct {
	IdProduct    uint       `gorm:"primaryKey" json:"idProduct"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	IdEnterprise int64      `json:"idEnterprise"`
	Enterprise   Enterprise `gorm:"foreignKey:IdEnterprise" json:"enterprise"`
}

type ProductOnSale struct {
	IdProductOnSale   int64     `gorm:"primaryKey;column:id_product_on_sale" json:"idProductOnSale"`
	Title             string    `json:"title"`
	IdProduct         int64     `json:"idProduct"`
	IdProductCatalog  int64     `json:"idProductCatalog"`
	Price             float64   `json:"price"`
	SaleStartDatetime time.Time `gorm:"column:sale_start_datetime" json:"saleStartDatetime"`
	SaleEndDatetime   time.Time `gorm:"column:sale_end_datetime" json:"saleEndDatetime"`
	Product           Product   `gorm:"foreignKey:IdProduct" json:"product"`
	Catalog           Catalog   `gorm:"foreignKey:IdProductCatalog" json:"catalog"`
}

func (Catalog) TableName() string {
	return "commerce.dim_product_catalog"
}

func (Enterprise) TableName() string {
	return "commerce.dim_enterprise"
}

func (Product) TableName() string {
	return "commerce.dim_product"
}

func (ProductOnSale) TableName() string {
	return "commerce.fact_product_on_sale"
}
