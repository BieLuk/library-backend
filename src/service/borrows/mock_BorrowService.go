// Code generated by mockery. DO NOT EDIT.

package borrows

import (
	dto "github.com/BieLuk/library-backend/src/dto"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockBorrowService is an autogenerated mock type for the BorrowService type
type MockBorrowService struct {
	mock.Mock
}

// CreateBorrow provides a mock function with given fields: request
func (_m *MockBorrowService) CreateBorrow(request dto.CreateBorrowRequest) (*dto.CreateBorrowResponse, error) {
	ret := _m.Called(request)

	var r0 *dto.CreateBorrowResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreateBorrowRequest) (*dto.CreateBorrowResponse, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(dto.CreateBorrowRequest) *dto.CreateBorrowResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CreateBorrowResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.CreateBorrowRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBookBorrowed provides a mock function with given fields: bookID
func (_m *MockBorrowService) IsBookBorrowed(bookID uuid.UUID) (*dto.IsBookBorrowedResponse, error) {
	ret := _m.Called(bookID)

	var r0 *dto.IsBookBorrowedResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*dto.IsBookBorrowedResponse, error)); ok {
		return rf(bookID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *dto.IsBookBorrowedResponse); ok {
		r0 = rf(bookID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.IsBookBorrowedResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(bookID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReturnBorrowedBook provides a mock function with given fields: bookID
func (_m *MockBorrowService) ReturnBorrowedBook(bookID uuid.UUID) error {
	ret := _m.Called(bookID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(bookID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockBorrowService creates a new instance of MockBorrowService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBorrowService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBorrowService {
	mock := &MockBorrowService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
