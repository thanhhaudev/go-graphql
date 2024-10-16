package service

import (
	"context"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type BookService struct {
	bookRepository repository.BookRepository
}

func (b *BookService) Create(ctx context.Context, input *model.CreateBookInput) (*model.Book, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	book := input.ToBook()
	if err := b.bookRepository.Create(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookService) FindAll(ctx context.Context) ([]*model.Book, error) {
	return b.bookRepository.FindAll(ctx)
}

func (b *BookService) FindByID(ctx context.Context, id int) (*model.Book, error) {
	return b.bookRepository.FindByID(ctx, id)
}

func (b *BookService) FindBooksByAuthorID(ctx context.Context, authorID int) ([]*model.Book, error) {
	return b.bookRepository.FindBooksByAuthorID(ctx, authorID)
}

func (b *BookService) FindBooksByBorrowerID(ctx context.Context, borrowerID int) ([]*model.Book, error) {
	return b.bookRepository.FindBooksByBorrowerID(ctx, borrowerID)
}

func NewBookService(bookRepository repository.BookRepository) *BookService {
	return &BookService{bookRepository: bookRepository}
}
