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

func TestBorrowerService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBorrowerRepo := mock_repository.NewMockBorrowerRepository(ctrl)
	mockBookRepo := mock_repository.NewMockBookRepository(ctrl)
	service := NewBorrowerService(mockBorrowerRepo, mockBookRepo)

	tests := []struct {
		name   string
		setup  func()
		run    func() (interface{}, error)
		verify func(t *testing.T, result interface{}, err error)
	}{
		{
			name: "CreateBorrower",
			setup: func() {
				mockBorrowerRepo.EXPECT().FindByTelNumber(gomock.Any(), "123456789").Return(nil, nil)
				mockBorrowerRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&model.Borrower{
					ID:        1,
					Name:      "John Doe",
					TelNumber: "123456789",
				}, nil)
			},
			run: func() (interface{}, error) {
				input := &model.CreateBorrowerInput{
					Name:      "John Doe",
					TelNumber: "123456789",
					Address:   "123 Main St",
					BirthDate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
				}
				return service.Create(context.Background(), input)
			},
			verify: func(t *testing.T, result interface{}, err error) {
				borrower := result.(*model.Borrower)
				assert.NoError(t, err)
				assert.NotNil(t, borrower)
				assert.Equal(t, "John Doe", borrower.Name)
				assert.Equal(t, "123456789", borrower.TelNumber)
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
