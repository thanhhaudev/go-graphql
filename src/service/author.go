package service

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/loader"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type AuthorService struct {
	authorRepository repository.AuthorRepository
}

func (a *AuthorService) Create(ctx context.Context, input *model.CreateAuthorInput) (*model.Author, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return a.authorRepository.Create(ctx, input.ToAuthor())
}

// GetAll returns all authors
func (a *AuthorService) GetAll(ctx context.Context) ([]*model.Author, error) {
	return a.authorRepository.GetAll(ctx)
}

// FindByID returns an author by id
func (a *AuthorService) FindByID(ctx context.Context, id int) (*model.Author, error) {
	return loader.FindAuthor(ctx, id)
}

func (a *AuthorService) GetByBookID(ctx context.Context, bookID int) ([]*model.Author, error) {
	book, err := loader.FindBook(ctx, bookID)
	if err != nil {
		return nil, err
	}

	return book.Authors, nil
}

func NewAuthorService(authorRepository repository.AuthorRepository) *AuthorService {
	return &AuthorService{authorRepository: authorRepository}
}
