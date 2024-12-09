package db

import (
	"database/sql"
	"encoding/json"

	"zocket_assignment/db/models"
)

// CreateUser inserts a new user into the database
func CreateUser(db *sql.DB, user models.User) (int, error) {
	var id int
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, user.Name, user.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// CreateProduct inserts a new product into the database
func CreateProduct(dbConn *sql.DB, productName, productDescription string, productImages, compressedProductImages []string, productPrice float64) (int, error) {
	query := `INSERT INTO products (product_name, product_description, product_images, compressed_product_images, product_price) 
              VALUES ($1, $2, $3, $4, $5) RETURNING id`

	// Prepare JSON arrays for the database
	var productID int
	err := dbConn.QueryRow(query, productName, productDescription, productImages, compressedProductImages, productPrice).Scan(&productID)
	if err != nil {
		return 0, err
	}
	return productID, nil
}

// GetProductByID retrieves a product by its ID
func GetProductByID(db *sql.DB, id int) (*models.Product, error) {
	query := `SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price, created_at, updated_at FROM products WHERE id = $1`
	row := db.QueryRow(query, id)

	var product models.Product
	var productImagesJSON, compressedProductImagesJSON []byte
	err := row.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription,
		&productImagesJSON, &compressedProductImagesJSON, &product.ProductPrice, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON fields
	json.Unmarshal(productImagesJSON, &product.ProductImages)
	json.Unmarshal(compressedProductImagesJSON, &product.CompressedProductImages)

	return &product, nil
}

// GetProducts retrieves all products for a specific user, with optional filtering
func GetProducts(db *sql.DB, userID int, minPrice, maxPrice float64, nameFilter string) ([]models.Product, error) {
	query := `SELECT id, user_id, product_name, product_description, product_images, compressed_product_images, product_price, created_at, updated_at 
              FROM products WHERE user_id = $1 AND product_price BETWEEN $2 AND $3 AND product_name ILIKE $4`

	rows, err := db.Query(query, userID, minPrice, maxPrice, "%"+nameFilter+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		var productImagesJSON, compressedProductImagesJSON []byte
		err := rows.Scan(&product.ID, &product.UserID, &product.ProductName, &product.ProductDescription,
			&productImagesJSON, &compressedProductImagesJSON, &product.ProductPrice, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}

		// Unmarshal JSON fields
		json.Unmarshal(productImagesJSON, &product.ProductImages)
		json.Unmarshal(compressedProductImagesJSON, &product.CompressedProductImages)

		products = append(products, product)
	}
	return products, nil
}
