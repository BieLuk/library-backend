package books

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/repository/books"
	"github.com/BieLuk/library-backend/src/repository/borrows"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateBook_ErrorBookExists(t *testing.T) {
	mockBookRepository := books.NewMockBooksRepository(t)
	mockBorrowsRepository := borrows.NewMockBorrowsRepository(t)
	mockBookRepository.Mock.On("CreateBook", &model.Book{
		Name:   "Test Name",
		Author: "Test Author",
		ISBN:   "Test ISBN",
	}).Return(nil, fmt.Errorf("book already exist")).Once()
	bookService := NewBookService(mockBookRepository, mockBorrowsRepository)

	request := dto.CreateBookRequest{
		Name:   "Test Name",
		Author: "Test Author",
		ISBN:   "Test ISBN",
	}
	_, err := bookService.CreateBook(request)
	assert.Error(t, err)

	mockBookRepository.AssertExpectations(t)
}

func TestCreateBook_Success(t *testing.T) {
	responseID := uuid.New()
	createdAt := time.Date(2023, 6, 14, 12, 36, 0, 0, time.UTC)
	updatedAt := time.Date(2023, 6, 14, 12, 36, 0, 0, time.UTC)
	mockBookRepository := books.NewMockBooksRepository(t)
	mockBorrowsRepository := borrows.NewMockBorrowsRepository(t)
	mockBookRepository.Mock.On("CreateBook", &model.Book{
		Name:   "Test Name",
		Author: "Test Author",
		ISBN:   "Test ISBN",
	}).Return(&model.Book{
		DBEntity: model.DBEntity{
			ID:        &responseID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		Name:   "Test Name",
		Author: "Test Author",
		ISBN:   "Test ISBN",
	}, nil).Once()
	bookService := NewBookService(mockBookRepository, mockBorrowsRepository)

	request := dto.CreateBookRequest{
		Name:   "Test Name",
		Author: "Test Author",
		ISBN:   "Test ISBN",
	}
	response, err := bookService.CreateBook(request)
	assert.NoError(t, err)

	assert.Equal(t, responseID, response.ID)

	mockBookRepository.AssertExpectations(t)
}
