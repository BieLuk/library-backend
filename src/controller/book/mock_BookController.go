// Code generated by mockery. DO NOT EDIT.

package book

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// MockBookController is an autogenerated mock type for the BookController type
type MockBookController struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: c
func (_m *MockBookController) CreateBook(c *gin.Context) {
	_m.Called(c)
}

// GetBook provides a mock function with given fields: c
func (_m *MockBookController) GetBook(c *gin.Context) {
	_m.Called(c)
}

// GetBooks provides a mock function with given fields: c
func (_m *MockBookController) GetBooks(c *gin.Context) {
	_m.Called(c)
}

// NewMockBookController creates a new instance of MockBookController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBookController(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBookController {
	mock := &MockBookController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
