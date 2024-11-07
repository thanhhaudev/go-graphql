package repository

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
)

type BookRepository interface {
	GetAll(ctx context.Context) ([]*model.Book, error)
	GetByIDs(ctx context.Context, ids []int) ([]*model.Book, error)
	FindByID(ctx context.Context, id int) (*model.Book, error)
	Create(ctx context.Context, model *model.Book) error
	Update(ctx context.Context, model *model.Book) error
	Delete(ctx context.Context, id int) error
}

type AuthorRepository interface {
	GetAll(ctx context.Context) ([]*model.Author, error)
	GetByIDs(ctx context.Context, ids []int) ([]*model.Author, error)
	FindByID(ctx context.Context, id int) (*model.Author, error)
	Create(ctx context.Context, model *model.Author) (*model.Author, error)
	Update(ctx context.Context, model *model.Author) (*model.Author, error)
	Delete(ctx context.Context, id int) error
}

type BorrowerRepository interface {
	GetAll(ctx context.Context) ([]*model.Borrower, error)
	GetByIDs(ctx context.Context, ids []int) ([]*model.Borrower, error)
	FindByID(ctx context.Context, id int) (*model.Borrower, error)
	FindByTelNumber(ctx context.Context, telNumber string) (*model.Borrower, error)
	Create(ctx context.Context, model *model.Borrower) (*model.Borrower, error)
	Update(ctx context.Context, model *model.Borrower) (*model.Borrower, error)
	BorrowBook(ctx context.Context, borrower *model.Borrower, book *model.Book) error
}
