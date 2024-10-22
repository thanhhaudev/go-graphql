package repository

import (
	"context"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
)

type BookRepository interface {
	FindAll(ctx context.Context) ([]*model.Book, error)
	FindByID(ctx context.Context, id int) (*model.Book, error)
	FindBooksByAuthorID(ctx context.Context, authorID int) ([]*model.Book, error)
	FindBooksByBorrowerID(ctx context.Context, borrowerID int) ([]*model.Book, error)
	Create(ctx context.Context, model *model.Book) error
	Update(ctx context.Context, model *model.Book) error
	Delete(ctx context.Context, id int) error
}

type AuthorRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	FindByID(ctx context.Context, id int) (*model.Author, error)
	FindByIDs(ctx context.Context, ids []int) ([]*model.Author, error)
	FindAuthorsByBookID(ctx context.Context, bookID int) ([]*model.Author, error)
	Create(ctx context.Context, model *model.Author) (*model.Author, error)
	Update(ctx context.Context, model *model.Author) (*model.Author, error)
	Delete(ctx context.Context, id int) error
}

type BorrowerRepository interface {
	FindAll(ctx context.Context) ([]*model.Borrower, error)
	FindByID(ctx context.Context, id int) (*model.Borrower, error)
	Create(ctx context.Context, input *model.BorrowerInput) error
	Update(ctx context.Context, input *model.BorrowerInput) error
}
