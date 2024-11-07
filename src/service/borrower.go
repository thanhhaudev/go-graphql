package service

import (
	"context"
	"fmt"
	"time"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/loader"
	"github.com/thanhhaudev/go-graphql/src/repository"
	"gorm.io/datatypes"
)

type BorrowerService struct {
	borrowerRepository repository.BorrowerRepository
	bookRepository     repository.BookRepository
}

func (b *BorrowerService) Create(ctx context.Context, input *model.CreateBorrowerInput) (*model.Borrower, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	if e, _ := b.borrowerRepository.FindByTelNumber(ctx, input.TelNumber); e != nil {
		return nil, fmt.Errorf("borrower with telNumber %q already exists", input.TelNumber)
	}

	return b.borrowerRepository.Create(ctx, input.ToBorrower())
}

func (b *BorrowerService) Update(ctx context.Context, id int, input *model.UpdateBorrowerInput) (*model.Borrower, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	borrower, err := b.borrowerRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	input.Patch(borrower)

	return b.borrowerRepository.Update(ctx, borrower)
}

func (b *BorrowerService) GetAll(ctx context.Context) ([]*model.Borrower, error) {
	return b.borrowerRepository.GetAll(ctx)
}

func (b *BorrowerService) FindByID(ctx context.Context, id int) (*model.Borrower, error) {
	return loader.FindBorrower(ctx, id)
}

func (b *BorrowerService) FindBooksByID(ctx context.Context, borrowerID int) ([]*model.Book, error) {
	borrower, err := loader.FindBorrower(ctx, borrowerID)
	if err != nil {
		return nil, err
	}

	books := make([]*model.Book, len(borrower.Borrowed))
	for i, bb := range borrower.Borrowed {
		books[i] = bb.Book
	}

	return books, nil
}

func (b *BorrowerService) BorrowBook(ctx context.Context, input *model.BorrowBookInput) (*model.Borrower, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	book, err := b.bookRepository.FindByID(ctx, input.BookID)
	if err != nil {
		return nil, err
	}

	// check if the book has enough quantity
	if book.Quantity < input.Quantity {
		return nil, fmt.Errorf("not enough quantity for book %q", book.Title)
	}

	book.DecrementQuantity(input.Quantity)

	borrower, err := b.borrowerRepository.FindByID(ctx, input.BorrowerID)
	if err != nil {
		return nil, err
	}

	// check if the borrower has already borrowed the book
	for _, bb := range borrower.Borrowed {
		if bb.BookID == input.BookID && bb.Status == model.BorrowedStatus {
			return nil, fmt.Errorf("borrower has already borrowed %d copies of book %q", bb.Quantity, book.Title)
		}
	}

	borrowed := model.BorrowerBook{
		Borrower: borrower,
		Book:     book,
		Quantity: input.Quantity,
		DueDate:  datatypes.Date(input.DueDate),
	}

	borrower.Borrowed = append(borrower.Borrowed, &borrowed)

	err = b.borrowerRepository.BorrowBook(ctx, borrower, book)
	if err != nil {
		return nil, err
	}

	return borrower, nil
}

func (b *BorrowerService) ReturnBook(ctx context.Context, input *model.ReturnBookInput) (*model.Borrower, error) {
	borrower, err := b.borrowerRepository.FindByID(ctx, input.BorrowerID)
	if err != nil {
		return nil, err
	}

	book, err := b.bookRepository.FindByID(ctx, input.BookID)
	if err != nil {
		return nil, err
	}

	// check if the borrower has borrowed the book
	var borrowed *model.BorrowerBook
	for _, bb := range borrower.Borrowed {
		if bb.BookID == book.ID && bb.Status == model.BorrowedStatus {
			borrowed = bb
			break
		}
	}

	if borrowed == nil {
		return nil, fmt.Errorf("borrower has not borrowed book %q", book.Title)
	}

	now := time.Now()
	book.IncrementQuantity(borrowed.Quantity)
	borrowed.Status = model.ReturnedStatus
	borrowed.ReturnedAt = &now

	err = b.borrowerRepository.BorrowBook(ctx, borrower, book)
	if err != nil {
		return nil, err
	}

	return borrower, nil
}

func NewBorrowerService(
	borrowerRepository repository.BorrowerRepository,
	bookRepository repository.BookRepository,
) *BorrowerService {
	return &BorrowerService{borrowerRepository: borrowerRepository, bookRepository: bookRepository}
}
