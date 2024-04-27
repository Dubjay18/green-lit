
# Green-Lit API

## Description

This is a RESTful API built with Go. It provides functionalities for user authentication, article management, and user management.

## Getting Started

### Prerequisites

- Go version 1.x
- PostgreSQL

### Installation

1. Clone the repository
```bash
git clone https://github.com/Dubjay18/green-lit.git
```
2. Navigate to the project directory
```bash
cd green-lit
```
3. Install the dependencies
```bash
go mod download
```
4. Set up your environment variables in a `.env` file. Refer to `env.copy` for required variables.

### Running the application

```bash
go run app.go
```

## API Endpoints

- `GET /`: Welcome endpoint
- `POST /auth-signIn`: User login
- `GET /users/{id}`: Get a user by ID
- `GET /articles/users/{id}`: Get articles by user ID
- `GET /users`: Get all users
- `GET /articles`: Get all articles
- `GET /articles/{id}`: Get an article by ID
- `POST /articles`: Create a new article

## Testing

To run the tests, use the following command:

```bash
go test ./...
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](LICENSE.md)
```

