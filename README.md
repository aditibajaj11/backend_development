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

System Architecture
The system follows a modular architecture with the following modules:
API Module: Handles all HTTP requests.
Database Module: Manages PostgreSQL interactions.
Message Queue Module: Interfaces with RabbitMQ.
Image Processing Microservice: Handles asynchronous image compression tasks.
Caching Module: Integrates with Redis for caching

