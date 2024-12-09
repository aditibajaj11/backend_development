package models

import "time"

// Product defines the structure of a product in the database
type Product struct {
	ID                      int       `json:"id,omitempty"`
	UserID                  int       `json:"user_id"`
	ProductName             string    `json:"product_name"`
	ProductDescription      string    `json:"product_description"`
	ProductImages           []string  `json:"product_images"`
	CompressedProductImages []string  `json:"compressed_product_images"`
	ProductPrice            float64   `json:"product_price"`
	CreatedAt               time.Time `json:"created_at,omitempty"`
	UpdatedAt               time.Time `json:"updated_at,omitempty"`
}
