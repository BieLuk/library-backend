package borrows

import (
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/repository/books"
	"github.com/BieLuk/library-backend/src/repository/borrows"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBorrowService_CreateBorrow(t *testing.T) {
	responseID := uuid.New()
	createdAt := time.Date(2023, 6, 14, 12, 36, 0, 0, time.UTC)
	updatedAt := createdAt
	mockBookRepository := books.NewMockBooksRepository(t)
	mockBorrowsRepository := borrows.NewMockBorrowsRepository(t)

	borrowEntity := &model.Borrow{
		BookID:    uuid.New(),
		TakenDate: createdAt,
	}

	mockBookRepository.Mock.
		On("GetBook", borrowEntity.BookID).
		Return(&model.Book{}, nil).Once()

	mockBorrowsRepository.Mock.
		On("CreateBorrow", borrowEntity).
		Return(&model.Borrow{
			DBEntity: model.DBEntity{
				ID:        &responseID,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
			BookID:    uuid.New(),
			TakenDate: createdAt,
		}, nil).Once()
	borrowService := NewBorrowService(mockBorrowsRepository, mockBookRepository)

	request := dto.CreateBorrowRequest{
		BookID:    borrowEntity.BookID,
		TakenDate: borrowEntity.TakenDate,
	}
	response, err := borrowService.CreateBorrow(request)
	assert.NoError(t, err)

	assert.Equal(t, responseID, response.ID)

	mockBookRepository.AssertExpectations(t)
}
