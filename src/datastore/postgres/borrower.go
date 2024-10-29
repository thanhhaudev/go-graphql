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

func (b *borrowerRepository) GetAll(ctx context.Context) ([]*model.Borrower, error) {
	var borrowers []*model.Borrower
	if err := b.db.WithContext(ctx).Find(&borrowers).Error; err != nil {
		return nil, err
	}

	return borrowers, nil
}

func (b *borrowerRepository) FindByID(ctx context.Context, id int) (*model.Borrower, error) {
	var borrower model.Borrower
	if err := b.db.WithContext(ctx).First(&borrower, id).Error; err != nil {
		return nil, err
	}

	return &borrower, nil
}

func (b *borrowerRepository) FindBorrowerBooksByID(ctx context.Context, borrowerID int) ([]*model.BorrowerBook, error) {
	var borrowerBooks []*model.BorrowerBook
	if err := b.db.WithContext(ctx).Preload("Book").Where("borrower_id = ?", borrowerID).Find(&borrowerBooks).Error; err != nil {
		return nil, err
	}

	return borrowerBooks, nil
}

func (b *borrowerRepository) FindByTelNumber(ctx context.Context, telNumber string) (*model.Borrower, error) {
	var borrower model.Borrower
	if err := b.db.WithContext(ctx).Where("tel_number = ?", telNumber).First(&borrower).Error; err != nil {
		return nil, err
	}

	return &borrower, nil
}

func (b *borrowerRepository) Create(ctx context.Context, input *model.Borrower) (*model.Borrower, error) {
	if err := b.db.WithContext(ctx).Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}

func (b *borrowerRepository) Update(ctx context.Context, input *model.Borrower) (*model.Borrower, error) {
	//TODO implement me
	panic("implement me")
}

func NewBorrowerRepository(db *gorm.DB) repository.BorrowerRepository {
	return &borrowerRepository{db: db}
}
