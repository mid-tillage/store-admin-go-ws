package models

import "time"

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
