package routes

import (
	"database/sql"
	"net/http"

	"zocket_assignment/api/handlers"

	"github.com/gorilla/mux"
)

// InitializeRoutes sets up the routes for the application
func InitializeRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	// Product routes with db connection passed to handlers
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Product API!"))
	}).Methods(http.MethodGet)

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateProduct(w, r, db)
	}).Methods(http.MethodPost)

	r.HandleFunc("/products/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProductByID(w, r, db)
	}).Methods(http.MethodGet)

	r.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetProducts(w, r, db)
	}).Methods(http.MethodGet)

	return r
}
