package resolver

import "github.com/thanhhaudev/go-graphql/src/service"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	authorService *service.AuthorService
}

func NewResolver(a *service.AuthorService) *Resolver {
	return &Resolver{
		authorService: a,
	}
}
