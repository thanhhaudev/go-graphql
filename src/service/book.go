package service

import (
	"context"
	"fmt"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type BookService struct {
	bookRepository   repository.BookRepository
	authorRepository repository.AuthorRepository
}

func (b *BookService) Create(ctx context.Context, input *model.CreateBookInput) (*model.Book, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	authors, err := b.authorRepository.FindByIDs(ctx, input.AuthorIds)
	if err != nil {
		return nil, err
	}

	if len(authors) != len(input.AuthorIds) {
		return nil, fmt.Errorf("some authors are not found")
	}

	book := input.ToBook()
	book.SetAuthors(authors)

	if err := b.bookRepository.Create(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookService) Update(ctx context.Context, id int, input *model.UpdateBookInput) (*model.Book, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	book, err := b.bookRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if input.AuthorIds != nil && len(input.AuthorIds) > 0 {
		authors, err := b.authorRepository.FindByIDs(ctx, input.AuthorIds)
		if err != nil {
			return nil, err
		}

		if len(authors) != len(input.AuthorIds) {
			return nil, fmt.Errorf("some authors are not found")
		}

		book.SetAuthors(authors)
	}

	input.Patch(book)

	if err := b.bookRepository.Update(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookService) GetAll(ctx context.Context) ([]*model.Book, error) {
	return b.bookRepository.GetAll(ctx)
}

func (b *BookService) FindByID(ctx context.Context, id int) (*model.Book, error) {
	return b.bookRepository.FindByID(ctx, id)
}

func (b *BookService) FindBooksByAuthorID(ctx context.Context, authorID int) ([]*model.Book, error) {
	return b.bookRepository.FindBooksByAuthorID(ctx, authorID)
}

func NewBookService(b repository.BookRepository, au repository.AuthorRepository) *BookService {
	return &BookService{bookRepository: b, authorRepository: au}
}
