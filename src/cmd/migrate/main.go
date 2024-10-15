package main

import (
	"github.com/joho/godotenv"
	"github.com/thanhhaudev/go-graphql/src/config"
	"github.com/thanhhaudev/go-graphql/src/migration"
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

	m := migration.Migrations(db.DB())
	if err := m.Migrate(); err != nil {
		log.Fatalf("could not run migration: %v", err)
	} else {
		log.Println("Migration ran successfully.")
	}
}
