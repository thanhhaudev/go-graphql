package loader

import (
	"context"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type authorLoader struct {
	authorRepository repository.AuthorRepository
}

// getAuthors implements a batch function that loads authors by IDs
func (a *authorLoader) getAuthors(ctx context.Context, ids []int) ([]*model.Author, []error) {
	authors, err := a.authorRepository.FindByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	return mapping(authors, ids), nil
}

func newAuthorLoader(authorRepository repository.AuthorRepository) *authorLoader {
	return &authorLoader{
		authorRepository: authorRepository,
	}
}
