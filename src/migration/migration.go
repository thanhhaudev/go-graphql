package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "00001_create_authors_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.Author{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("authors")
			},
		},
		{
			ID: "00002_create_books_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.Book{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("books")
			},
		},
		{
			ID: "00003_create_authors_books_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.AuthorBook{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("authors_books")
			},
		},
		{
			ID: "00004_create_borrowers_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.Borrower{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("borrowers")
			},
		},
		{
			ID: "00005_create_borrowers_books_table",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.BorrowersBooks{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("borrowers_books")
			},
		},
	})
}
