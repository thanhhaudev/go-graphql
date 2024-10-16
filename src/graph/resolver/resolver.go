package resolver

import "github.com/thanhhaudev/go-graphql/src/service"

//go:generate go run github.com/99designs/gqlgen generate

// Resolver contains application dependencies

type Resolver struct {
	authorService   *service.AuthorService
	bookService     *service.BookService
	borrowerService *service.BorrowerService
}

func NewResolver(a *service.AuthorService, b *service.BookService, br *service.BorrowerService) *Resolver {
	return &Resolver{
		authorService:   a,
		bookService:     b,
		borrowerService: br,
	}
}
