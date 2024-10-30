package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thanhhaudev/go-graphql/src/graph/model"
	mock_repository "github.com/thanhhaudev/go-graphql/src/mock"
	"go.uber.org/mock/gomock"
)

func TestAuthorService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorRepo := mock_repository.NewMockAuthorRepository(ctrl)
	service := NewAuthorService(mockAuthorRepo)

	tests := []struct {
		name   string
		setup  func()
		run    func() (interface{}, error)
		verify func(t *testing.T, result interface{}, err error)
	}{
		{
			name: "CreateAuthor",
			setup: func() {
				mockAuthorRepo.EXPECT().Create(context.Background(), gomock.Any()).Return(&model.Author{
					ID:   1,
					Name: "John Doe",
				}, nil)

			},
			run: func() (interface{}, error) {
				return service.Create(context.Background(), &model.CreateAuthorInput{
					Name: "John Doe",
				})
			},
			verify: func(t *testing.T, result interface{}, err error) {
				author := result.(*model.Author)
				assert.NoError(t, err)
				assert.NotNil(t, author)
				assert.Equal(t, "John Doe", author.Name)
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
