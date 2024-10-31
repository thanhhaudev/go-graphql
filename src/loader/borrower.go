package loader

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
)

type borrowerLoader struct {
	borrowerRepository repository.BorrowerRepository
}

// getBorrowers implements a batch function that loads borrowers by IDs
func (b *borrowerLoader) getBorrowers(ctx context.Context, ids []int) ([]*model.Borrower, []error) {
	borrowers, err := b.borrowerRepository.FindByIDs(ctx, ids)
	if err != nil {
		return nil, []error{err}
	}

	return mapping(borrowers, ids), nil
}

func newBorrowerLoader(borrowerRepository repository.BorrowerRepository) *borrowerLoader {
	return &borrowerLoader{
		borrowerRepository: borrowerRepository,
	}
}
