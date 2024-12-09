module zocket_assignment

go 1.23

require github.com/lib/pq v1.10.9 // PostgreSQL driver for Go

require github.com/gorilla/mux v1.8.1

require github.com/joho/godotenv v1.5.1 // indirect

replace github.com/lib/pq => github.com/lib/pq v1.10.9 // Ensure the correct version of pq is used
