package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"strconv"

	"github.com/thanhhaudev/go-graphql/src/graph/generated"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input model.CreateBookInput) (*model.Book, error) {
	return r.bookService.Create(ctx, &input)
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id string, input model.UpdateBookInput) (*model.Book, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return r.bookService.Update(ctx, intID, &input)
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.CreateAuthorInput) (*model.Author, error) {
	return r.authorService.Create(ctx, &input)
}

// CreateBorrower is the resolver for the createBorrower field.
func (r *mutationResolver) CreateBorrower(ctx context.Context, input model.CreateBorrowerInput) (*model.Borrower, error) {
	return r.borrowerService.Create(ctx, &input)
}

// UpdateBorrower is the resolver for the updateBorrower field.
func (r *mutationResolver) UpdateBorrower(ctx context.Context, id string, input model.UpdateBorrowerInput) (*model.Borrower, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return r.borrowerService.Update(ctx, intID, &input)
}

// BorrowBook is the resolver for the borrowBook field.
func (r *mutationResolver) BorrowBook(ctx context.Context, input *model.BorrowBookInput) (*model.Borrower, error) {
	return r.borrowerService.BorrowBook(ctx, input)
}

// ReturnBook is the resolver for the returnBook field.
func (r *mutationResolver) ReturnBook(ctx context.Context, input *model.ReturnBookInput) (*model.Borrower, error) {
	return r.borrowerService.ReturnBook(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
