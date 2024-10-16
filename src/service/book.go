package service

import (
	"context"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type BookService struct {
	bookRepository repository.BookRepository
}

func (b *BookService) FindAll(ctx context.Context) ([]*model.Book, error) {
	return b.bookRepository.FindAll(ctx)
}

func (b *BookService) FindByID(ctx context.Context, id int) (*model.Book, error) {
	return b.bookRepository.FindByID(ctx, id)
}

func NewBookService(bookRepository repository.BookRepository) *BookService {
	return &BookService{bookRepository: bookRepository}
}
