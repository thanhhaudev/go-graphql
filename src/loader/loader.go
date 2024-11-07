package loader

import (
	"context"
	"net/http"
	"time"

	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"github.com/thanhhaudev/go-graphql/src/repository"
	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

type Model interface {
	GetID() int
}

const (
	loadersKey = ctxKey("dataLoaders")
)

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	AuthorLoader   *dataloadgen.Loader[int, *model.Author]
	BookLoader     *dataloadgen.Loader[int, *model.Book]
	BorrowerLoader *dataloadgen.Loader[int, *model.Borrower]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(
	authorRepository repository.AuthorRepository,
	bookRepository repository.BookRepository,
	borrowerRepository repository.BorrowerRepository,
) *Loaders {
	al := newAuthorLoader(authorRepository)
	bl := newBookLoader(bookRepository)
	brl := newBorrowerLoader(borrowerRepository)
	w := dataloadgen.WithWait(10 * time.Millisecond)

	return &Loaders{
		AuthorLoader:   dataloadgen.NewLoader(al.getAuthors, w),
		BookLoader:     dataloadgen.NewLoader(bl.getBooks, w),
		BorrowerLoader: dataloadgen.NewLoader(brl.getBorrowers, w),
	}
}

// Middleware injects data loaders into the context
// This middleware likely injects data loaders into the request context, optimizing data fetching.
// This accessibly allows the resolvers to load data in a batch, reducing the number of queries to the database.
func Middleware(
	authorRepository repository.AuthorRepository,
	bookRepository repository.BookRepository,
	borrowerRepository repository.BorrowerRepository,
	next http.Handler,
) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(authorRepository, bookRepository, borrowerRepository)
		r = r.WithContext(context.WithValue(r.Context(), loadersKey, loader))

		next.ServeHTTP(w, r)
	})
}

// mapping maps the models to the same order as the input IDs
func mapping[M Model](models []M, ids []int) []M {
	// Create a map to easily find authors by ID
	idMap := make(map[int]M)
	for _, m := range models {
		idMap[m.GetID()] = m
	}

	// Prepare the result slice in the same order as the input IDs
	// This is important because the result will be used in the same order as the input IDs
	result := make([]M, len(ids))
	for i, id := range ids {
		result[i] = idMap[id]
	}

	return result
}

// For returns the data loader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// FindAuthor loads an author by ID
func FindAuthor(ctx context.Context, id int) (*model.Author, error) {
	return For(ctx).AuthorLoader.Load(ctx, id)
}

// GetAuthors loads authors by IDs
func GetAuthors(ctx context.Context, ids []int) ([]*model.Author, error) {
	return For(ctx).AuthorLoader.LoadAll(ctx, ids)
}

func FindBook(ctx context.Context, id int) (*model.Book, error) {
	return For(ctx).BookLoader.Load(ctx, id)
}

func GetBooks(ctx context.Context, ids []int) ([]*model.Book, error) {
	return For(ctx).BookLoader.LoadAll(ctx, ids)
}

func FindBorrower(ctx context.Context, id int) (*model.Borrower, error) {
	return For(ctx).BorrowerLoader.Load(ctx, id)
}

func GetBorrowers(ctx context.Context, ids []int) ([]*model.Borrower, error) {
	return For(ctx).BorrowerLoader.LoadAll(ctx, ids)
}
