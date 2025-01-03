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

func (a *authorRepository) GetAll(ctx context.Context) ([]*model.Author, error) {
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

func (a *authorRepository) GetByIDs(ctx context.Context, ids []int) ([]*model.Author, error) {
	var authors []*model.Author
	if err := a.db.WithContext(ctx).Preload("Books", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC")
	}).Where("id IN (?)", ids).
		Order("created_at DESC").
		Find(&authors).Error; err != nil {
		return nil, err
	}

	return authors, nil
}

func (a *authorRepository) Create(ctx context.Context, model *model.Author) (*model.Author, error) {
	if err := a.db.WithContext(ctx).Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (a *authorRepository) Update(ctx context.Context, model *model.Author) (*model.Author, error) {
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
