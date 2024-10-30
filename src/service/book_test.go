package service

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
	mock_repository "github.com/thanhhaudev/go-graphql/src/mock"
	"go.uber.org/mock/gomock"
)

func TestBookService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookRepo := mock_repository.NewMockBookRepository(ctrl)
	mockAuthorRepo := mock_repository.NewMockAuthorRepository(ctrl)
	service := NewBookService(mockBookRepo, mockAuthorRepo)

	tests := []struct {
		name   string
		setup  func()
		run    func() (interface{}, error)
		verify func(t *testing.T, result interface{}, err error)
	}{
		{
			name: "CreateBook",
			setup: func() {
				mockBookRepo.EXPECT().Create(context.Background(), gomock.Any()).Return(nil)
				mockAuthorRepo.EXPECT().FindByIDs(context.Background(), []int{1}).Return([]*model.Author{
					{
						ID: 1,
					},
				}, nil)
			},
			run: func() (interface{}, error) {
				return service.Create(context.Background(), &model.CreateBookInput{
					Title:     "Animal Farm",
					Quantity:  1,
					Rating:    5,
					AuthorIds: []int{1},
					PublishAt: time.Date(1945, time.August, 17, 0, 0, 0, 0, time.UTC),
				})
			},
			verify: func(t *testing.T, result interface{}, err error) {
				book := result.(*model.Book)
				assert.NoError(t, err)
				assert.NotNil(t, book)
				assert.Equal(t, "Animal Farm", book.Title)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			result, err := tt.run()
			tt.verify(t, result, err)
		})
	}
}
