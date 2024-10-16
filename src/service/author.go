package service

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type AuthorService struct {
	authorRepository repository.AuthorRepository
}

// FindAll returns all authors
func (a *AuthorService) FindAll(ctx context.Context) ([]*model.Author, error) {
	return a.authorRepository.FindAll(ctx)
}

// FindByID returns an author by id
func (a *AuthorService) FindByID(ctx context.Context, id int) (*model.Author, error) {
	return a.authorRepository.FindByID(ctx, id)
}

func (a *AuthorService) FindAuthorsByBookID(ctx context.Context, bookID int) ([]*model.Author, error) {
	return a.authorRepository.FindAuthorsByBookID(ctx, bookID)
}

func NewAuthorService(authorRepository repository.AuthorRepository) *AuthorService {
	return &AuthorService{authorRepository: authorRepository}
}
