# Mini Auth

A simple Go mini authentication system built with Fiber and MySQL.

## What is implemented so far

- HTTP server using `github.com/gofiber/fiber/v3`
- MySQL database connection via `database/sql` and `github.com/go-sql-driver/mysql`
- Environment configuration loading with `github.com/joho/godotenv`
- Basic health/status endpoint at `/`
- User creation endpoint at `/add`
- DB connection checker utility
- Shared types for `USER` and `DBSTATUS`

## Project structure

- `main.go` — application entry point and route definitions
- `db/db.go` — database connection setup and pool configuration
- `utils/utils.go` — database health check helper
- `types/types.go` — shared domain types

## API Endpoints

### GET `/`

Returns a JSON status response. Example:

{
  "status": {
    "Success": true,
    "Message": "Connected"
  },
  "server": "mini_auth"
}

### POST `/add`

Adds a new user record to the `users` table.

Expected JSON body:

{
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "password123"
}

Response on success:

{
  "message": "Success",
  "name": "Jane Doe",
  "email": "jane@example.com",
  "password": "password123"
}

## Environment variables

The application expects a `.env` file with the following values:

- `DB_HOST`
- `DB_USER`
- `DB_PASS`
- `DB_NAME`

## How to run

1. Install Go 1.26 or newer.
2. Create a `.env` file with MySQL credentials.
3. Run `go mod tidy` to download dependencies.
4. Start the server with `go run main.go`.
5. The app listens on port `3000`.

## Notes

- Passwords are currently stored directly as provided; password hashing is not implemented yet.
- There is no authentication, validation, or error handling beyond basic request parsing and insert failure handling.
- The `users` table must exist in the configured MySQL database.
- The project is currently focused on a minimal API for adding users and checking DB connectivity.
