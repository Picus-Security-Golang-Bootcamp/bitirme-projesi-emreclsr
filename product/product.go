package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	Type     string  `json:"type"`
	IsDelete bool    `json:"is_delete"`
}
