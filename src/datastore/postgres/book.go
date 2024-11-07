package postgres

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func (b *bookRepository) GetByIDs(ctx context.Context, ids []int) ([]*model.Book, error) {
	var books []*model.Book
	if err := b.db.WithContext(ctx).Preload("Authors").Find(&books, ids).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (b *bookRepository) GetAll(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	if err := b.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (b *bookRepository) FindByID(ctx context.Context, id int) (*model.Book, error) {
	var book model.Book
	if err := b.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (b *bookRepository) Create(ctx context.Context, model *model.Book) error {
	if err := b.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}

	return nil
}

func (b *bookRepository) Update(ctx context.Context, model *model.Book) error {
	// Start a transaction
	tx := b.db.WithContext(ctx).Begin()

	// Update the book details
	if err := tx.Omit("Authors").Updates(model).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update the Authors association
	if err := tx.Model(model).Association("Authors").Replace(model.Authors); err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

func (b *bookRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{db: db}
}
