package handlers

import (
	"database/sql"
	"encoding/json"

	"log"
	"net/http"
	"strconv"

	"zocket_assignment/db"
	"zocket_assignment/db/models"

	"github.com/gorilla/mux"
)

// CreateProduct handles POST /products
// CreateProduct handles POST /products
func CreateProduct(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	// Validate Content-Type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	// Read and decode the request body
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid JSON body: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate required fields
	if product.ProductName == "" || len(product.ProductImages) == 0 || product.ProductPrice <= 0 {
		http.Error(w, "Missing or invalid required fields", http.StatusBadRequest)
		return
	}

	// Insert the product into the database
	productID, err := db.CreateProduct(dbConn, product.ProductName, product.ProductDescription, product.ProductImages, product.CompressedProductImages, product.ProductPrice)
	if err != nil {
		log.Println("Error creating product in DB:", err)
		http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the created product ID
	response := map[string]int{"product_id": productID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetProductByID handles GET /products/{id}
func GetProductByID(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	// Parse the product ID from the URL
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the product by ID
	product, err := db.GetProductByID(dbConn, productID)
	if err != nil {
		log.Println("Error retrieving product:", err)
		http.Error(w, "Product not found: "+err.Error(), http.StatusNotFound)
		return
	}

	// Respond with the product details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GetProducts handles GET /products
func GetProducts(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	// Parse query parameters
	userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
	minPrice, _ := strconv.ParseFloat(r.URL.Query().Get("min_price"), 64)
	maxPrice, _ := strconv.ParseFloat(r.URL.Query().Get("max_price"), 64)
	nameFilter := r.URL.Query().Get("name_filter")

	// Retrieve the products
	products, err := db.GetProducts(dbConn, userID, minPrice, maxPrice, nameFilter)
	if err != nil {
		log.Println("Error retrieving products:", err)
		http.Error(w, "Failed to retrieve products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the product list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
