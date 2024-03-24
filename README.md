# Book Store Golang Project
This is a Golang project for a book store, designed for learning purposes. The project implements basic CRUD (Create, Read, Update, Delete) operations for managing books.

## Features

- Book Management:
  - [x] Add new books to the store.
  - [x] Retrieve a list of all books available in the store.
  - [x] Get detailed information about a specific book by its ID.
  - [x] Delete books from the store.
  - [ ] Get detailed information about a specific book by its ISBN.
  - [ ] Update existing books with new information or modifications.

- Author Management:
  - [x] Add new authors to the store.
  - [x] Retrieve a list of all authors.
  - [x] Get detailed information about a specific author by their ID.
  - [x] Delete authors from the store.
  - [ ] Update existing author information.

- User Management:
  - [x] Create new users with unique API keys.
  - [x] Authenticate users with middleware authentication.
  - [x] Manage user permissions and access control.

- Middleware Authentication:
  - [x] Secure API endpoints using middleware authentication.
  - [x] Authenticate user requests using API keys.
  - [x] Ensure only authorized users have access to sensitive endpoints.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- Golang installed on your system
- PostgreSQL database server installed and running
- Goose CLI tool installed (for database migrations)
- SQLC CLI tool installed (for generating database queries)
- Git installed on your system

## Installation

To install and run this project, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/Hayoun01/book_store_api
    ```
2. Navigate to the project directory:

    ```bash
    cd book-store
    ```
3. Install project dependencies:

    ```bash
    go mod tidy
    ```
4. Set up the PostgreSQL database:

    * Create a new database named `book_store`.
    * Import the database schema from the `schema.sql` file in the `db` directory.

5. Configure environment variables:

    * Copy the .env.example file to .env.
    * Update the .env file with your PostgreSQL database credentials and app Port.

6. Run the application:

    ```bash
    go run ./cmd/main/
    ```
The application should now be running. Access the API at http://localhost:8080.

## Usage
Once the application is running, you can interact with the API endpoints using tools like cURL, Postman, or a web browser:

- **GET `/books`**: Retrieve a list of all books.
- **POST `/create_book`**: Add a new book to the store (authentication required).
- **GET `/books/{book_id}`**: Get details of a specific book by its ID.
- **DELETE `/books/{book_id}`**: Delete a book from the store (authentication required).
- **POST `/create_author`**: Create a new author (authentication required).
- **GET `/authors`**: Retrieve a list of all authors.
- **POST `/users`**: Create a new user (no authentication required).
- **GET `/ready`**: Check if the server is ready (no authentication required).

## Project Hierarchy
The project follows the following directory structure:
- `cmd/main`: Main entry point of the application.
- `pkg/auth`: Contains authentication-related code.
- `pkg/config`: Configuration settings for the application.
- `pkg/controllers`: Controllers for handling HTTP requests.
- `pkg/db`: Database-related code, including migrations, queries, and SQLC generated code.
- `pkg/models`: Data models used by the application.
- `pkg/routes`: Route definitions for the API endpoints.
- `pkg/utils`: Utility functions used throughout the application.

## Database Migrations
Database migrations are managed using Goose. Use the following Makefile commands to manage migrations:

```bash
# Apply pending migrations.
make migrateup
# Rollback the last migration.
make migratedown
```
## SQLC Queries
SQLC is used to generate type-safe database queries. Use the following Makefile command to generate SQLC queries:
```bash
# Generate SQLC queries.
make sqlc 
```
## Contributing
Contributions are welcome! Here's how you can contribute to this project:

1. Fork the repository.
2. Create a new branch for your feature or bug fix: git checkout -b feature-name.
3. Make your changes and commit them: git commit -m 'Add new feature'.
4. Push to your fork: git push origin feature-name.
5. Create a pull request on the original repository.

## License
This project is licensed under the MIT License. See the LICENSE file for details.

Feel free to customize this template according to your project's specific requirements and structure.