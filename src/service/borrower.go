package service

import (
	"context"
	"fmt"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type BorrowerService struct {
	borrowerRepository repository.BorrowerRepository
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
	return b.borrowerRepository.FindByID(ctx, id)
}

func (b *BorrowerService) FindBooksByID(ctx context.Context, borrowerID int) ([]*model.Book, error) {
	borrowed, err := b.borrowerRepository.FindBorrowerBooksByID(ctx, borrowerID)
	if err != nil {
		return nil, err
	}

	books := make([]*model.Book, 0, len(borrowed))
	for _, bb := range borrowed {
		books = append(books, bb.Book)
	}

	return books, nil
}

func NewBorrowerService(borrowerRepository repository.BorrowerRepository) *BorrowerService {
	return &BorrowerService{borrowerRepository: borrowerRepository}
}
