# Simple GraphQL Library Management System

This project is a **Simple Library Management System** built with **Golang** and **gqlgen**. It provides a GraphQL API for managing books, authors, and borrowers, along with borrowing and returning books.

## Features

- **Book Management**: Create, update, and retrieve book information.
- **Author Management**: Manage author records.
- **Borrower Management**: Handle borrower data.
- **Borrowing System**: Track book borrowing and returning status.

## Database Schema

The system's database schema is as follows:

```plaintext
+--------------------+         +--------------------+        +--------------------+
|      authors       |         |     authors_books  |        |       books        |
+--------------------+         +--------------------+        +--------------------+
| id                 |<------->| author_id          |<------>| id                 |
| name               |         | book_id            |        | title              |
| created_at         |         |                    |        | description        |
| updated_at         |         +--------------------+        | quantity           |
| deleted_at         |                                       | rating             |
+--------------------+                                       | publish_at         |
                                                             | created_at         |
                                                             | updated_at         |
                                                             | deleted_at         |
                                                             +--------------------+
                                                                       |
                                                                       |
                                                                       |
                                                             +---------v---------+
                                                             |   borrower_books   |
+--------------------+                                       +--------------------+
|     borrowers      |                                       | id                 |
+--------------------+                                       | borrower_id        |
| id                 |<------------------------------------->| book_id            |
| name               |                                       | borrowed_at        |
| address            |                                       | returned_at        |
| tel_number         |                                       | due_date           |
| birth_date         |                                       | status             |
| created_at         |                                       | quantity           |
| updated_at         |                                       | created_at         |
| deleted_at         |                                       | updated_at         |
+--------------------+                                       | deleted_at         |
                                                             +--------------------+
```

## Technologies Used

- **Go**: Primary programming language for the backend.
- **GORM**: Object-Relational Mapping (ORM) library for Go.
- **gqlgen**: Library for generating GraphQL server code in Go.

## Getting Started

### Prerequisites

| Requirement                                                | Version       |
|------------------------------------------------------------|---------------|
| [Docker Compose](https://docs.docker.com/compose/install/) | Latest        |

### Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/thanhhaudev/go-graphql.git
    cd go-graphql
    ```

2. **Build and start the Docker containers**:
    ```sh
    make build
    make run
    ```

3. **Migrate the database schema**:
    ```sh
    make migrate
    ```

4. **Seed the database with sample data**:
    ```sh
    make seed
    ```

### Running the Application

1. Ensure the Docker containers are running:
    ```sh
    make run && make logs
    ```

2. Access the GraphQL playground at `http://localhost:8080/`.

3. To stop the Docker containers:
    ```sh
    make stop
    ```

## Usage Examples

### GraphQL Queries and Mutations
- **Get all Books**:
    ```graphql
    query {
        books {
            id
            title
            authors {
                id
                name
            }
        }
    }
    ```

- **Get a Book by ID**:
    ```graphql
    query {
        book(id: 1) {
            id
            title
            authors {
                id
                name
            }
        }
    }
    ```

- **Get all Authors**:
    ```graphql
    query {
        authors {
            id
            name
            books {
                id
                title
            }
        }
    }
    ```

- **Get all Borrowers**:
    ```graphql
    query {
        borrowers {
            id
            name
            books {
                id
                title
            }
        }
    }
    ```

- **Create a Book**:
    ```graphql
    mutation {
        createBook(
            input: {title: "Title", authorIds: [1, 2], publishAt: "2024-09-30T07:00:00.000Z", quantity: 100, rating: 4.2}
        ) {
            id
            title
            createdAt
        }
    }
    ```

- **Borrow a Book**:
    ```graphql
    mutation {
        borrowBook(
            input: {bookId: 1, borrowerId: 1, quantity: 5, dueDate: "2024-09-30T07:00:00.000Z"}
        ) {
            name
            books {
                title
            }
        }
    }
    ```

- **Return a Book**:
    ```graphql
    mutation {
        returnBook(input: {bookId: 2, borrowerId: 1}) {
            name
            books {
                id
                title
            }
        }
    }
    ```

## Contributing

Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and test them.
4. Submit a pull request.
