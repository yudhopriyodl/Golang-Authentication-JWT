# Golang Authentication with JWT

This project implements a simple authentication system using Golang and JSON Web Tokens (JWT). It provides endpoints for user registration, login, and accessing protected resources.

## Features

*   User Registration
*   User Login
*   JWT-based Authentication
*   Protected Routes

## Installation

To get started with this project, follow these steps:

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/golang-authentication-jwt.git
    cd golang-authentication-jwt
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Set up environment variables:**
    Create a `.env` file in the root directory and add the following:

    ```
    JWT_SECRET=your_super_secret_jwt_key
    PORT=8080
    ```
    Replace `your_super_secret_jwt_key` with a strong, random secret key.

## Usage

1.  **Run the application:**

    ```bash
    go run main.go
    ```

    The application will start on the port specified in your `.env` file (default: `8080`).

2.  **API Endpoints:**

    *   **Register User:**
        `POST /register`
        Body:
        ```json
        {
            "username": "testuser",
            "password": "password123"
        }
        ```

    *   **Login User:**
        `POST /login`
        Body:
        ```json
        {
            "username": "testuser",
            "password": "password123"
        }
        ```
        Response will include a JWT token.

    *   **Access Protected Route:**
        `GET /protected`
        Headers:
        `Authorization: Bearer <your_jwt_token>`

## Project Structure

```
.
├── .env
├── go.mod
├── go.sum
├── main.go
├── controllers/
│   └── auth_controller.go
├── middlewares/
│   └── jwt_middleware.go
├── models/
│   └── user.go
└── utils/
    └── jwt.go
