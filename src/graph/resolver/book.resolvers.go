package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"github.com/thanhhaudev/go-graphql/src/graph/generated"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
)

// Authors is the resolver for the authors field.
func (r *bookResolver) Authors(ctx context.Context, obj *model.Book) ([]*model.Author, error) {
	authors, err := r.authorService.FindAuthorsByBookID(ctx, obj.ID)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

type bookResolver struct{ *Resolver }
