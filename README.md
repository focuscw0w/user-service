# User Microservice (API in Go)

## Overview
This project is a **user management microservice** that exposes a **REST API** for authentication and user management.  
It is written in Go, uses **SQLite** as the database, and implements **JWT-based authentication**.  
The service is designed to be part of a microservices architecture.

## Features
### Authentication
- `POST /sign-up` → Register a new user
- `POST /sign-in` → Log in (sets `jwt` cookie)
- `POST /sign-out` → Log out (clears the cookie)

### Users
- `GET /users` → Get all users
- `GET /users/{id}` → Get user by ID *(requires login)*
- `PUT /users/update/{id}` → Update user *(requires login)*
- `DELETE /users/{id}` → Delete user *(requires login)*

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/user-service.git
   cd user-service
2. Create a .env file with configuration:
   SECRET_KEY=supersecret
   PORT=8080
3. go run main.go
4. http://localhost:8080

## Example Requests
### 1. Sign Up
Register a new user:
  ```bash
  curl -X POST http://localhost:8080/sign-up \
    -H "Content-Type: application/json" \
    -d '{"username":"student","email":"student@example.com","password":"pass123"}'
  ```

### 2. Sign in
Log in and save the authentication cookie (cookies.txt):
  ```bash
  curl -i -c cookies.txt -X POST http://localhost:8080/sign-in \
    -H "Content-Type: application/json" \
    -d '{"email":"student@example.com","password":"pass123"}'
  ```

### 3. Access Protected Route
Use the saved cookie to access a protected endpoint:
  ```
  curl -b cookies.txt http://localhost:8080/users/1
  ```

### 4. Update User
Update user info (requires login):
  ```bash
  curl -b cookies.txt -X PUT http://localhost:8080/users/update/1 \
    -H "Content-Type: application/json" \
    -d '{"username":"newName","email":"newmail@example.com","password":"newpass"}'
  ```

### 5. Delete User
Delete a user (requires login):
  ```bash
  curl -b cookies.txt -X DELETE http://localhost:8080/users/1
  ```

### 6. Sign out
Log out (clears the cookie):
  ```bash
  curl -b cookies.txt -X POST http://localhost:8080/sign-out
  ```


