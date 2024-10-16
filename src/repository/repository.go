package repository

import (
	"context"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
)

type BookRepository interface {
	FindAll(ctx context.Context) ([]*model.Book, error)
	FindByID(ctx context.Context, id int) (*model.Book, error)
	FindBooksByAuthorID(ctx context.Context, authorID int) ([]*model.Book, error)
	Create(ctx context.Context, input *model.BookInput) error
	Update(ctx context.Context, input *model.BookInput) error
	Delete(ctx context.Context, id int) error
}

type AuthorRepository interface {
	FindAll(ctx context.Context) ([]*model.Author, error)
	FindByID(ctx context.Context, id int) (*model.Author, error)
	Create(ctx context.Context, input *model.AuthorInput) error
	Update(ctx context.Context, input *model.AuthorInput) error
	Delete(ctx context.Context, id int) error
}

type BorrowerRepository interface {
	FindAll(ctx context.Context) ([]*model.Borrower, error)
	FindByID(ctx context.Context, id int) (*model.Borrower, error)
	Create(ctx context.Context, input *model.BorrowerInput) error
	Update(ctx context.Context, input *model.BorrowerInput) error
}
