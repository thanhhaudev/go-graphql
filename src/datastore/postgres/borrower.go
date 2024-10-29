package postgres

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := b.db.WithContext(ctx).Preload("Borrowed").First(&borrower, id).Error; err != nil {
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

func (b *borrowerRepository) Create(ctx context.Context, model *model.Borrower) (*model.Borrower, error) {
	if err := b.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (b *borrowerRepository) Update(ctx context.Context, model *model.Borrower) (*model.Borrower, error) {
	if err := b.db.WithContext(ctx).Model(model).Updates(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (b *borrowerRepository) BorrowBook(ctx context.Context, borrower *model.Borrower, book *model.Book) error {
	return b.db.Transaction(func(tx *gorm.DB) error {
		// update borrower
		if err := b.db.WithContext(ctx).Model(borrower).
			Clauses(clause.OnConflict{DoNothing: true}).
			Session(&gorm.Session{FullSaveAssociations: true}).
			Updates(borrower).Error; err != nil {
			return err
		}

		// update book quantity
		if err := b.db.WithContext(ctx).Model(book).Update("quantity", book.Quantity).Error; err != nil {
			return err
		}

		return nil
	})
}

func NewBorrowerRepository(db *gorm.DB) repository.BorrowerRepository {
	return &borrowerRepository{db: db}
}
