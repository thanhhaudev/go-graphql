package loader

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type bookLoader struct {
	bookRepository repository.BookRepository
}

// getBooks implements a batch function that loads books by IDs
func (b *bookLoader) getBooks(ctx context.Context, ids []int) ([]*model.Book, []error) {
	books, err := b.bookRepository.GetByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	return mapping(books, ids), nil
}

func newBookLoader(bookRepository repository.BookRepository) *bookLoader {
	return &bookLoader{
		bookRepository: bookRepository,
	}
}
