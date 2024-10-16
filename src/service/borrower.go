package service

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type BorrowerService struct {
	borrowerRepository repository.BorrowerRepository
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
