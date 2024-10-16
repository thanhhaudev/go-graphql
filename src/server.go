package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/thanhhaudev/go-graphql/src/config"
	"github.com/thanhhaudev/go-graphql/src/datastore/postgres"
	"github.com/thanhhaudev/go-graphql/src/graph/generated"
	"github.com/thanhhaudev/go-graphql/src/graph/resolver"
	"github.com/thanhhaudev/go-graphql/src/service"
)

const defaultPort = "8080"

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get port from environment variables or use default port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Connect to database
	db, err := config.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	authorRepo := postgres.NewAuthorRepository(db.DB())
	bookRepo := postgres.NewBookRepository(db.DB())
	borrowerRepo := postgres.NewBorrowerRepository(db.DB())

	authorService := service.NewAuthorService(authorRepo)
	bookService := service.NewBookService(bookRepo)
	borrowerService := service.NewBorrowerService(borrowerRepo)

	defer db.Close() // Close database connection when main function exits

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver.NewResolver(authorService, bookService, borrowerService),
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
