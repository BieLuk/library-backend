// Code generated by mockery. DO NOT EDIT.

package borrows

import (
	model "github.com/BieLuk/library-backend/src/model"
	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// MockBorrowsRepository is an autogenerated mock type for the BorrowsRepository type
type MockBorrowsRepository struct {
	mock.Mock
}

// CreateBorrow provides a mock function with given fields: book
func (_m *MockBorrowsRepository) CreateBorrow(book *model.Borrow) (*model.Borrow, error) {
	ret := _m.Called(book)

	var r0 *model.Borrow
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.Borrow) (*model.Borrow, error)); ok {
		return rf(book)
	}
	if rf, ok := ret.Get(0).(func(*model.Borrow) *model.Borrow); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Borrow)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.Borrow) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBorrowsByBookID provides a mock function with given fields: bookID
func (_m *MockBorrowsRepository) GetBorrowsByBookID(bookID uuid.UUID) ([]*model.Borrow, error) {
	ret := _m.Called(bookID)

	var r0 []*model.Borrow
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*model.Borrow, error)); ok {
		return rf(bookID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*model.Borrow); ok {
		r0 = rf(bookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Borrow)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBorrowsNotBroughtByBookID provides a mock function with given fields: bookID
func (_m *MockBorrowsRepository) GetBorrowsNotBroughtByBookID(bookID uuid.UUID) ([]*model.Borrow, error) {
	ret := _m.Called(bookID)

	var r0 []*model.Borrow
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*model.Borrow, error)); ok {
		return rf(bookID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*model.Borrow); ok {
		r0 = rf(bookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Borrow)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBorrowBroughtDateByBookID provides a mock function with given fields: bookID, broughtDate
func (_m *MockBorrowsRepository) UpdateBorrowBroughtDateByBookID(bookID uuid.UUID, broughtDate time.Time) error {
	ret := _m.Called(bookID, broughtDate)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, time.Time) error); ok {
		r0 = rf(bookID, broughtDate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockBorrowsRepository creates a new instance of MockBorrowsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBorrowsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBorrowsRepository {
	mock := &MockBorrowsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
