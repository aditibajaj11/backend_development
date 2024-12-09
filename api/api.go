package api

import (
	"database/sql"
	"zocket_assignment/api/routes"

	"github.com/gorilla/mux"
)

// SetRouter sets up and returns the router with routes initialized
func SetRouter(db *sql.DB) *mux.Router {
	return routes.InitializeRoutes(db)
}
