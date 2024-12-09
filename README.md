This repository contains the source code for a backend system built in Go for a product management application. The project emphasizes architectural best practices, including asynchronous processing, caching, logging, and high scalability.

Project Overview
The system is a RESTful API designed for managing products with features like:
Asynchronous image processing.
Redis caching for optimized performance.
Enhanced structured logging.
PostgreSQL for robust data storage

Features
API Endpoints:
POST /products: Add product details.
GET /products/:id: Retrieve product details, leveraging caching for performance.
GET /products: Fetch all products for a specific user, with optional filters.
Data Storage:
PostgreSQL for storing users and products, including processed image data.

