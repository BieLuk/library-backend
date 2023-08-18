package borrows

import (
	"errors"
	"fmt"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/repository/borrows"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"time"
)

type BorrowService interface {
	CreateBorrow(request dto.CreateBorrowRequest) (*dto.CreateBorrowResponse, error)
	IsBookBorrowed(bookID uuid.UUID) (*dto.IsBookBorrowedResponse, error)
	ReturnBorrowedBook(bookID uuid.UUID) error
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
	borrowEntity, err := s.borrowRepository.CreateBorrow(borrowEntity)
	if err != nil {
		return nil, fmt.Errorf("cannot create borrowEntity: %w", err)
	}

	return &dto.CreateBorrowResponse{ID: *borrowEntity.ID}, nil
}

func (s *borrowService) IsBookBorrowed(bookID uuid.UUID) (*dto.IsBookBorrowedResponse, error) {
	_, err := s.borrowRepository.GetBorrowsByBookID(bookID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, mongo.ErrNoDocuments) {
			return &dto.IsBookBorrowedResponse{IsBorrowed: false}, nil
		}
		return nil, fmt.Errorf("cannot get borrow: %w", err)
	}

	return &dto.IsBookBorrowedResponse{IsBorrowed: true}, nil
}

func (s *borrowService) ReturnBorrowedBook(bookID uuid.UUID) error {
	err := s.borrowRepository.UpdateBorrowBroughtDateByBookID(bookID, time.Now())
	if err != nil {
		return fmt.Errorf("cannot return borrowed book: %w", err)
	}

	return nil
}
