package postgres

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
	"gorm.io/gorm"
)

type borrowerRepository struct {
	db *gorm.DB
}

func (b borrowerRepository) FindAll(ctx context.Context) ([]*model.Borrower, error) {
	var borrowers []*model.Borrower
	if err := b.db.WithContext(ctx).Find(&borrowers).Error; err != nil {
		return nil, err
	}

	return borrowers, nil
}

func (b borrowerRepository) FindByID(ctx context.Context, id int) (*model.Borrower, error) {
	var borrower model.Borrower
	if err := b.db.WithContext(ctx).First(&borrower, id).Error; err != nil {
		return nil, err
	}

	return &borrower, nil
}

func (b borrowerRepository) Create(ctx context.Context, input *model.BorrowerInput) error {
	//TODO implement me
	panic("implement me")
}

func (b borrowerRepository) Update(ctx context.Context, input *model.BorrowerInput) error {
	//TODO implement me
	panic("implement me")
}

func NewBorrowerRepository(db *gorm.DB) repository.BorrowerRepository {
	return &borrowerRepository{db: db}
}
