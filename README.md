# User-Service-CRUD-Assignment-Go-Gin-GORM

## Author

**Nziza Samuel**  
Email: nzizataylora@gmail.com  

## User Service API

A User Management Service API built with **Go** and **Gin**, providing functionality to manage users, including creating, reading, updating, and deleting accounts. It features **role-based access control** (admin-only endpoints) and **JWT authentication**.

## Features

- User registration and login with JWT authentication
- Role-based access control (admin vs user)
- CRUD operations on users:
  - List all users (admin only)
  - Get user by ID
  - Update user details
  - Delete user (admin only)
- Swagger documentation available for easy API exploration

## Technologies

- Go 
- Gin (HTTP web framework)
- JWT for authentication
- PostgreSQL 
- Swagger for API documentation

## Installation

1. **Clone the repository:**
```bash
git clone https://github.com/Nziza21/user-service.git
cd user-service


Install dependencies:

go mod tidy

Run the server:

go run main.go


The API will be accessible at: http://localhost:8080

API Documentation

Swagger UI is available at:
http://localhost:8080/docs/index.html
