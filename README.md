# Library App

## Getting Started

To get started with using the Library App, follow the steps below:

### Prerequisites

- Go programming language installed on your system.
- PostgreSQL database installed and running.

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Build the project:

   ```
   go build
   ```

### Running the Server

Run the server using the following command:

```
go run .
```

## API Endpoints

The following endpoints are available in the API:

### Users

- `POST /api/v1/register`: Register a new user.
- `GET /api/v1/users`: Get all users.
- `GET /api/v1/users/{userId}`: Get a user by ID.
- `PUT /api/v1/users/{userId}`: Update a user by ID.
- `DELETE /api/v1/users/{userId}`: Delete a user by ID.

## DB structure

```
CREATE TABLE IF NOT EXISTS users (
    id           bigserial PRIMARY KEY,
    createdAt    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updatedAt    timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    username     text                        NOT NULL,
    password     text                        NOT NULL
);
```
---
