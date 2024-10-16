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

func (b bookRepository) FindAll(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	if err := b.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (b bookRepository) FindByID(ctx context.Context, id int) (*model.Book, error) {
	var book model.Book
	if err := b.db.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func (b bookRepository) Create(ctx context.Context, input *model.BookInput) error {
	//TODO implement me
	panic("implement me")
}

func (b bookRepository) Update(ctx context.Context, input *model.BookInput) error {
	//TODO implement me
	panic("implement me")
}

func (b bookRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{db: db}
}