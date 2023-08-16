package borrows

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/db/postgres"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/google/uuid"
)

type borrowsPostgresRepository struct {
}

func NewBorrowPostgresRepository() *borrowsPostgresRepository {
	return &borrowsPostgresRepository{}
}

func (r *borrowsPostgresRepository) CreateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	if result := postgres.GetDB().Create(borrow); result.Error != nil {
		return nil, fmt.Errorf("error creating borrow in database: %w", result.Error)
	}
	return borrow, nil
}

func (r *borrowsPostgresRepository) GetBorrowByBookID(bookID uuid.UUID) (*model.Borrow, error) {
	var borrow *model.Borrow
	if result := postgres.GetDB().Model(&model.Borrow{}).Joins("Book").Where("book_id = ?", bookID).Take(&borrow); result.Error != nil {
		return nil, fmt.Errorf("error retrieving borrow from database: %w", result.Error)
	}
	return borrow, nil
}
