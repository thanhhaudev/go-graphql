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

func (b *BorrowerService) FindAll(ctx context.Context) ([]*model.Borrower, error) {
	return b.borrowerRepository.FindAll(ctx)
}

func (b *BorrowerService) FindByID(ctx context.Context, id int) (*model.Borrower, error) {
	return b.borrowerRepository.FindByID(ctx, id)
}

func NewBorrowerService(borrowerRepository repository.BorrowerRepository) *BorrowerService {
	return &BorrowerService{borrowerRepository: borrowerRepository}
}
