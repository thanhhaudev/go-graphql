package resolver

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
}

func NewResolver() *Resolver {
	return &Resolver{}
}
