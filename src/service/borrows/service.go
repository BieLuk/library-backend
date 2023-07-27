package borrows

import (
	"errors"
	"fmt"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/repository/borrows"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BorrowService interface {
	CreateBorrow(request dto.CreateBorrowRequest) (*dto.CreateBorrowResponse, error)
	IsBookBorrowed(bookID uuid.UUID) (*dto.IsBookBorrowedResponse, error)
}

type borrowService struct {
	borrowRepository borrows.BorrowsRepository
}

func NewBorrowService(borrowRepository borrows.BorrowsRepository) *borrowService {
	return &borrowService{
		borrowRepository: borrowRepository,
	}
}

func (s *borrowService) CreateBorrow(request dto.CreateBorrowRequest) (*dto.CreateBorrowResponse, error) {
	borrowEntity := &model.Borrow{
		BookID:    request.BookID,
		TakenDate: request.TakenDate,
	}
	if err := s.borrowRepository.CreateBorrow(borrowEntity); err != nil {
		return nil, fmt.Errorf("cannot create borrowEntity: %w", err)
	}

	return &dto.CreateBorrowResponse{ID: *borrowEntity.ID}, nil
}

func (s *borrowService) IsBookBorrowed(bookID uuid.UUID) (*dto.IsBookBorrowedResponse, error) {
	_, err := s.borrowRepository.GetBorrowByBookID(bookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &dto.IsBookBorrowedResponse{IsBorrowed: false}, nil
		}
		return nil, fmt.Errorf("cannot get borrow: %w", err)
	}

	return &dto.IsBookBorrowedResponse{IsBorrowed: true}, nil
}
