package postgres

import (
	"context"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func (a *authorRepository) FindAll(ctx context.Context) ([]*model.Author, error) {
	var authors []*model.Author
	if err := a.db.WithContext(ctx).Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (a *authorRepository) FindByID(ctx context.Context, id int) (*model.Author, error) {
	var author model.Author
	if err := a.db.WithContext(ctx).First(&author, id).Error; err != nil {
		return nil, err
	}

	return &author, nil
}

func (a *authorRepository) FindByIDs(ctx context.Context, ids []int) ([]*model.Author, error) {
	var authors []*model.Author
	if err := a.db.WithContext(ctx).Where("id IN (?)", ids).Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (a *authorRepository) FindAuthorsByBookID(ctx context.Context, bookID int) ([]*model.Author, error) {
	var book model.Book
	if err := a.db.WithContext(ctx).Preload("Authors").First(&book, bookID).Error; err != nil {
		return nil, err
	}

	return book.Authors, nil
}

func (a *authorRepository) Create(ctx context.Context, input *model.AuthorInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepository) Update(ctx context.Context, input *model.AuthorInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *authorRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthorRepository(db *gorm.DB) repository.AuthorRepository {
	return &authorRepository{db: db}
}
