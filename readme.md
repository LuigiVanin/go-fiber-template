# Go Boilerplate API

A simple REST API boilerplate built with Go, Fiber, and PostgreSQL.

## Features

- User authentication (Sign Up / Sign In)
- PostgreSQL database with GORM
- Request validation
- Error handling with Problem Details (RFC 7807)
- CORS enabled

## Tech Stack

- **Framework**: Fiber v2
- **Database**: PostgreSQL
- **ORM**: GORM
- **Validation**: go-playground/validator
- **Environment**: godotenv

## Prerequisites

- Go 1.24.3 or higher
- PostgreSQL
- Make (optional)

## Getting Started

### 1. Clone the repository

```bash
git clone <repository-url>
cd boilerplate
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Set up environment variables

Create a `.env` file in the root directory:

```bash
cp .env.example .env
```

Edit `.env` with your database credentials:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=your_database
DB_SSLMODE=disable

PORT=3000

APP_ENV=development
JWT_SERCRET=SUPER
```

### 4. Run the application

Using Make:
```bash
make run
```

Or directly with Go:
```bash
go run cmd/main.go
```

For development with hot reload (requires [Air](https://github.com/cosmtrek/air)):
```bash
make dev
```

### 5. Build the application

```bash
make build
```

This creates an executable in `./bin/main`

## Project Structure

```
boilerplate/
├── cmd/
│   └── main.go                 # Application entry point
├── app/
│   ├── bootstrap.go            # App initialization
│   ├── common/                 # Common interfaces and utilities
│   ├── middleware/             # HTTP middlewares
│   ├── models/dto/             # Data Transfer Objects
│   └── module/                 # Feature modules
│       ├── auth/               # Authentication module
│       │   ├── controller/
│       │   └── service/
│       └── user/               # User module
│           └── repository/
├── internal/
│   ├── configuration/          # Environment configuration
│   ├── database/               # Database setup and migrations
│   │   └── entity/             # Database entities
│   └── errors/                 # Error handling
├── utils/                      # Utility functions
├── .env.example               # Example environment variables
├── Makefile                   # Build commands
├── go.mod                     # Go modules
└── README.md                  # This file
```

## Error Handling

The API uses RFC 7807 Problem Details for HTTP APIs. Error responses follow this format:

```json
{
  "type": "about:blank",
  "title": "Validation error",
  "status": 422,
  "detail": "Request validation failed",
  "instance": "/signup",
  "errors": [
    {
      "field": "Email",
      "tag": "email",
      "param": "",
      "value": "invalid-email"
    }
  ]
}
```

## Database Migration

Database migrations run automatically on application startup. The `users` table will be created if it doesn't exist.

**User Schema:**
- `id` (uint, primary key)
- `name` (text, required)
- `email` (text, unique, required)
- `password` (text, required)
- `created_at` (timestamp)
- `updated_at` (timestamp)

To wireup new entities to the migration it is necessary to change the Migrate method on the [infra/database/client.go](infra/database/client.go) file.

## Development

### Available Make Commands

```bash
make run    # Run the application
make dev    # Run with hot reload (requires Air)
make build  # Build the application
```

### Adding New Modules

1. Create a new folder under `app/module/`
2. Implement controller, service, and repository layers
3. Register routes in the controller
4. Wire dependencies in `cmd/main.go`

## License

See [license.md](license.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.