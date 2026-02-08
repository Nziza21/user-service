# User-Service-CRUD-Assignment-Go-Gin-GORM

## Author

**Nziza Samuel**  
Email: samuel.nziza@example.com  

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


Configure environment variables (.env file):

DB_DSN=postgres://username:password@localhost:5432/user_service_db
PORT=8800
JWT_SECRET=mysecretpassword


Run the server:

go run cmd/server/main.go


The API will be accessible at: http://localhost:8800

API Documentation

Swagger UI is available at:
http://localhost:8800/docs/index.html

Example Requests
Create a User
curl -X POST http://localhost:8800/api/v1/users \
-H "Content-Type: application/json" \
-d '{
  "fullName": "Nziza S. Samuel",
  "email": "nziza@example.com",
  "phone": "+250788999888",
  "role": "user",
  "status": "active"
}'

Get User by ID
curl http://localhost:8800/api/v1/users/<user-id>

Update User
curl -X PATCH http://localhost:8800/api/v1/users/<user-id> \
-H "Content-Type: application/json" \
-d '{
  "fullName": "Nziza Samuel",
  "phone": "+250787946474"
}'

Delete User (Admin Only)
curl -X DELETE http://localhost:8800/api/v1/users/<user-id> \
-H "Authorization: Bearer <admin-jwt-token>"

User Login
curl -X POST http://localhost:8800/api/v1/auth/login \
-H "Content-Type: application/json" \
-d '{
  "email": "nziza@example.com",
  "password": "yourpassword"
}'

Reset Password
curl -X POST http://localhost:8800/api/v1/auth/reset-password \
-H "Content-Type: application/json" \
-d '{
  "email": "nziza@example.com",
  "otp": "generated otp",
  "new_password": "newpassword123"
}'

License

MIT License. See LICENSE