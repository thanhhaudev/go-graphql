package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"github.com/thanhhaudev/go-graphql/src/config"
	"github.com/thanhhaudev/go-graphql/src/migration"
)

func main() {
	// Define command-line flags
	action := flag.String("action", "migrate", "Action to perform: migrate or rollback")
	flag.Parse()

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
	switch *action {
	case "migrate":
		if err := m.Migrate(); err != nil {
			log.Fatalf("could not run migration: %v", err)
		} else {
			log.Println("Migration ran successfully.")
		}
	case "rollback":
		if err := m.RollbackLast(); err != nil {
			log.Fatalf("could not rollback migration: %v", err)
		} else {
			log.Println("Rollback ran successfully.")
		}
	case "refresh":
		if err := m.RollbackTo("00001_create_authors_table"); err != nil {
			log.Fatalf("could not rollback migration: %v", err)
		} else {
			if err := m.RollbackLast(); err != nil { // Rollback the last migration, because this library doesn't support rollback all migrations
				log.Fatalf("could not rollback migration: %v", err)
			}

			log.Println("Rollback ran successfully.")

			if err := m.Migrate(); err != nil {
				log.Fatalf("could not run migration: %v", err)
			} else {
				log.Println("Migration ran successfully.")
			}
		}
	default:
		log.Fatalf("Unknown action: %s", *action)
	}
}
