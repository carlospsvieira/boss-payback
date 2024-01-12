# Boss Payback

## Overview

Boss Payback is a web application built with [Fiber](https://github.com/gofiber/fiber), designed for managing users, roles, expenses, and workflows. The system utilizes JWT for authentication and connects to a PostgreSQL database.

## Prerequisites

- Go version 1.21.5
- Docker
- Docker Compose

## Dependencies

### Go Modules

The application manages dependencies using Go modules. Key dependencies include:

```go
github.com/gofiber/fiber/v2 v2.51.0
github.com/golang-jwt/jwt/v5 v5.2.0
github.com/google/uuid v1.4.0
github.com/joho/godotenv v1.5.1
gorm.io/driver/postgres v1.5.4
gorm.io/gorm v1.25.5
```

## Setup

### Environment Variables

Ensure you have a `.env` file in the root directory with necessary environment variables:

```env
# Database Configuration
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_HOST=
DB_PORT=

# Authentication
JWT_SECRET_KEY=
ADMIN_USERNAME=
ADMIN_PASSWORD=

# Uploads Directory Path
UPLOADS_DIR_PATH=./example_uploads/
```
## Docker

To deploy and run the application using Docker, make use of the provided `docker-compose.yml` configuration:

```yaml
version: '3.8'

services:
  postgres-db:
    image: postgres:latest
    ports:
      - "5432:5432"
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: ${DB_NAME}
    volumes:
      - postgres-data:/var/lib/postgresql/data

  fiber-app:
    build: 
      context: .
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    environment:
      - DB_HOST=postgres-db
      - DB_USER=${DB_USER}
      - DB_PASSWORD=${DB_PASSWORD}
      - DB_NAME=${DB_NAME}
      - JWT_SECRET_KEY=${JWT_SECRET_KEY}
      - ADMIN_USERNAME=${ADMIN_USERNAME}
      - ADMIN_PASSWORD=${ADMIN_PASSWORD}
      - UPLOADS_DIR_PATH=${UPLOADS_DIR_PATH}
    depends_on:
      - postgres-db

volumes:
  postgres-data:
```
To initiate the Docker containers, execute:

```bash
docker-compose up -d
```

## Development

For a streamlined development experience, consider using [Air](https://github.com/cosmtrek/air) for auto-reloading. After setting up Docker containers, you can initiate Air with:

```bash
go get -u github.com/cosmtrek/air
air
```

## Makefile

The project includes a `Makefile` to simplify frequent development tasks. Available targets include:

- `build`: Compile the application.
- `run`: Build and execute the application.
- `clean`: Eliminate build artifacts.
- `deps`: Install necessary dependencies.
- `fmt`: Format the codebase.
- `help`: View available make commands.

Example command:

```bash
make run
```
## Documentation

For detailed insights into the API routes, models, and handlers, refer to the [Documentation](https://github.com/carlospsvieira/boss-payback/docs/api.md)

