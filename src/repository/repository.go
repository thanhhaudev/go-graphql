package repository

import "github.com/thanhhaudev/go-graphql/src/graph/model"

type BookRepository interface {
	FindAll() ([]*model.Book, error)
	FindByID(id int) (*model.Book, error)
	Create(input *model.BookInput) error
	Update(input *model.BookInput) error
	Delete(id int) error
}

type AuthorRepository interface {
	FindAll() ([]*model.Author, error)
	FindByID(id int) (*model.Author, error)
	Create(input *model.AuthorInput) error
	Update(input *model.AuthorInput) error
	Delete(id int) error
}

type BorrowerRepository interface {
	FindAll() ([]*model.Borrower, error)
	FindByID(id int) (*model.Borrower, error)
	Create(input *model.BorrowerInput) error
	Update(input *model.BorrowerInput) error
}
