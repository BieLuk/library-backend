package borrows

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/db"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/google/uuid"
)

type BorrowsRepository interface {
	CreateBorrow(book *model.Borrow) error
	GetBorrowByBookID(bookID uuid.UUID) (*model.Borrow, error)
}

type borrowsRepository struct {
}

func NewBorrowRepository() *borrowsRepository {
	return &borrowsRepository{}
}

func (r *borrowsRepository) CreateBorrow(borrow *model.Borrow) error {
	if result := db.GetDB().Create(borrow); result.Error != nil {
		return fmt.Errorf("error creating borrow in database: %w", result.Error)
	}
	return nil
}

func (r *borrowsRepository) GetBorrowByBookID(bookID uuid.UUID) (*model.Borrow, error) {
	var borrow *model.Borrow
	if result := db.GetDB().Model(&model.Borrow{}).Joins("Book").Where("book_id = ?", bookID).Take(&borrow); result.Error != nil {
		return nil, fmt.Errorf("error retrieving borrow from database: %w", result.Error)
	}
	return borrow, nil
}
