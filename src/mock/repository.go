// Code generated by MockGen. DO NOT EDIT.
// Source: repository/repository.go
//
// Generated by this command:
//
//	mockgen -source repository/repository.go -destination mock/repository.go
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	model "github.com/thanhhaudev/go-graphql/src/graph/model"
	gomock "go.uber.org/mock/gomock"
)

// MockBookRepository is a mock of BookRepository interface.
type MockBookRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBookRepositoryMockRecorder
	isgomock struct{}
}

// MockBookRepositoryMockRecorder is the mock recorder for MockBookRepository.
type MockBookRepositoryMockRecorder struct {
	mock *MockBookRepository
}

// NewMockBookRepository creates a new mock instance.
func NewMockBookRepository(ctrl *gomock.Controller) *MockBookRepository {
	mock := &MockBookRepository{ctrl: ctrl}
	mock.recorder = &MockBookRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookRepository) EXPECT() *MockBookRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockBookRepository) Create(ctx context.Context, m *model.Book) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBookRepositoryMockRecorder) Create(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBookRepository)(nil).Create), ctx, m)
}

// Delete mocks base method.
func (m *MockBookRepository) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBookRepositoryMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBookRepository)(nil).Delete), ctx, id)
}

// FindByID mocks base method.
func (m *MockBookRepository) FindByID(ctx context.Context, id int) (*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockBookRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockBookRepository)(nil).FindByID), ctx, id)
}

// GetAll mocks base method.
func (m *MockBookRepository) GetAll(ctx context.Context) ([]*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBookRepositoryMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBookRepository)(nil).GetAll), ctx)
}

// GetByIDs mocks base method.
func (m *MockBookRepository) GetByIDs(ctx context.Context, ids []int) ([]*model.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, ids)
	ret0, _ := ret[0].([]*model.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockBookRepositoryMockRecorder) GetByIDs(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockBookRepository)(nil).GetByIDs), ctx, ids)
}

// Update mocks base method.
func (m_2 *MockBookRepository) Update(ctx context.Context, m *model.Book) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBookRepositoryMockRecorder) Update(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBookRepository)(nil).Update), ctx, m)
}

// MockAuthorRepository is a mock of AuthorRepository interface.
type MockAuthorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorRepositoryMockRecorder
	isgomock struct{}
}

// MockAuthorRepositoryMockRecorder is the mock recorder for MockAuthorRepository.
type MockAuthorRepositoryMockRecorder struct {
	mock *MockAuthorRepository
}

// NewMockAuthorRepository creates a new mock instance.
func NewMockAuthorRepository(ctrl *gomock.Controller) *MockAuthorRepository {
	mock := &MockAuthorRepository{ctrl: ctrl}
	mock.recorder = &MockAuthorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorRepository) EXPECT() *MockAuthorRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m_2 *MockAuthorRepository) Create(ctx context.Context, m *model.Author) (*model.Author, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthorRepositoryMockRecorder) Create(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthorRepository)(nil).Create), ctx, m)
}

// Delete mocks base method.
func (m *MockAuthorRepository) Delete(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorRepositoryMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthorRepository)(nil).Delete), ctx, id)
}

// FindByID mocks base method.
func (m *MockAuthorRepository) FindByID(ctx context.Context, id int) (*model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockAuthorRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockAuthorRepository)(nil).FindByID), ctx, id)
}

// GetAll mocks base method.
func (m *MockAuthorRepository) GetAll(ctx context.Context) ([]*model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAuthorRepositoryMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAuthorRepository)(nil).GetAll), ctx)
}

// GetByIDs mocks base method.
func (m *MockAuthorRepository) GetByIDs(ctx context.Context, ids []int) ([]*model.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, ids)
	ret0, _ := ret[0].([]*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockAuthorRepositoryMockRecorder) GetByIDs(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockAuthorRepository)(nil).GetByIDs), ctx, ids)
}

// Update mocks base method.
func (m_2 *MockAuthorRepository) Update(ctx context.Context, m *model.Author) (*model.Author, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(*model.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAuthorRepositoryMockRecorder) Update(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthorRepository)(nil).Update), ctx, m)
}

// MockBorrowerRepository is a mock of BorrowerRepository interface.
type MockBorrowerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBorrowerRepositoryMockRecorder
	isgomock struct{}
}

// MockBorrowerRepositoryMockRecorder is the mock recorder for MockBorrowerRepository.
type MockBorrowerRepositoryMockRecorder struct {
	mock *MockBorrowerRepository
}

// NewMockBorrowerRepository creates a new mock instance.
func NewMockBorrowerRepository(ctrl *gomock.Controller) *MockBorrowerRepository {
	mock := &MockBorrowerRepository{ctrl: ctrl}
	mock.recorder = &MockBorrowerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBorrowerRepository) EXPECT() *MockBorrowerRepositoryMockRecorder {
	return m.recorder
}

// BorrowBook mocks base method.
func (m *MockBorrowerRepository) BorrowBook(ctx context.Context, borrower *model.Borrower, book *model.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BorrowBook", ctx, borrower, book)
	ret0, _ := ret[0].(error)
	return ret0
}

// BorrowBook indicates an expected call of BorrowBook.
func (mr *MockBorrowerRepositoryMockRecorder) BorrowBook(ctx, borrower, book any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BorrowBook", reflect.TypeOf((*MockBorrowerRepository)(nil).BorrowBook), ctx, borrower, book)
}

// Create mocks base method.
func (m_2 *MockBorrowerRepository) Create(ctx context.Context, m *model.Borrower) (*model.Borrower, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Create", ctx, m)
	ret0, _ := ret[0].(*model.Borrower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockBorrowerRepositoryMockRecorder) Create(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBorrowerRepository)(nil).Create), ctx, m)
}

// FindByID mocks base method.
func (m *MockBorrowerRepository) FindByID(ctx context.Context, id int) (*model.Borrower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*model.Borrower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockBorrowerRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockBorrowerRepository)(nil).FindByID), ctx, id)
}

// FindByTelNumber mocks base method.
func (m *MockBorrowerRepository) FindByTelNumber(ctx context.Context, telNumber string) (*model.Borrower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByTelNumber", ctx, telNumber)
	ret0, _ := ret[0].(*model.Borrower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByTelNumber indicates an expected call of FindByTelNumber.
func (mr *MockBorrowerRepositoryMockRecorder) FindByTelNumber(ctx, telNumber any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByTelNumber", reflect.TypeOf((*MockBorrowerRepository)(nil).FindByTelNumber), ctx, telNumber)
}

// GetAll mocks base method.
func (m *MockBorrowerRepository) GetAll(ctx context.Context) ([]*model.Borrower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*model.Borrower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBorrowerRepositoryMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBorrowerRepository)(nil).GetAll), ctx)
}

// GetByIDs mocks base method.
func (m *MockBorrowerRepository) GetByIDs(ctx context.Context, ids []int) ([]*model.Borrower, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIDs", ctx, ids)
	ret0, _ := ret[0].([]*model.Borrower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs.
func (mr *MockBorrowerRepositoryMockRecorder) GetByIDs(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIDs", reflect.TypeOf((*MockBorrowerRepository)(nil).GetByIDs), ctx, ids)
}

// Update mocks base method.
func (m_2 *MockBorrowerRepository) Update(ctx context.Context, m *model.Borrower) (*model.Borrower, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(*model.Borrower)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockBorrowerRepositoryMockRecorder) Update(ctx, m any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBorrowerRepository)(nil).Update), ctx, m)
}
