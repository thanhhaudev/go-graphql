package main

import (
	"github.com/joho/godotenv"
	"github.com/thanhhaudev/go-graphql/src/config"
	"github.com/thanhhaudev/go-graphql/src/seeder"
	"log"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := config.NewDB()
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	defer db.Close()

	seeder.SeedAuthors(db.DB(), seeder.Authors)
	seeder.SeedBooks(db.DB(), seeder.Books)
	seeder.SeedBorrowers(db.DB(), seeder.Borrowers)
}
