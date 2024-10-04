# Golang Todo-Api 

## Description
The Golang Todo-Api is a robust and efficient RESTful API built using Go (Golang) and the Gin web framework. This project provides a backend service for managing todo tasks, allowing users to create, read, update, and delete todo items. It features user authentication using JWT tokens, ensuring secure access to personal todo lists.

Key features include:
- RESTful API endpoints for todo CRUD operations
- User authentication and authorization
- JWT-based secure routes
- Database integration for persistent storage
- Middleware for request logging and error handling
- Scalable and modular codebase structure

This API serves as an excellent backend for todo list applications, providing a solid foundation that can be easily extended or integrated with various frontend technologies.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Installation

1. Clone the repository:
   ``
   git clone https://github.com/Muyi2905/TodoApi
   ```
2. Navigate to the project directory:
   ```
   cd into repository
   ```
3. Install dependencies:
   ```
   go mod tidy
   ```
4. Set up your environment variables 

## Usage

To run the application:
go run cmd/main.go


## API Endpoints


- `GET /api/v1/users`: Retrieve all users
- `POST /api/v1/users`: Create a new user
- `GET /api/v1/users/:id`: Retrieve a specific user
- `PUT /api/v1/users/:id`: Update a user
- `DELETE /api/v1/users/:id`: Delete a user




## Environment Variables

The following environment variables are required:

- `JWT_SECRET`: Secret key for JWT token generation and validation
- `DB_CONNECTION_STRING`: Database connection string

## Project Structure

```
.
├── cmd/
│   └── main.go
├── controllers/
├── middleware/
├── models/
├── routes/
├── .env
├── .gitignore
├── go.mod
└── go.sum
```

Brief description of each directory:
- `cmd/`: Contains the main application entry point
- `controllers/`: Handles the application logic
- `middleware/`: Custom middleware functions
- `models/`: Data models and database interactions
- `routes/`: Defines API routes

## Contributing

Instructions for how to contribute to your project. For example:

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

https://roadmap.sh/projects/todo-list-api