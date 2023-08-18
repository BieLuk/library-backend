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

func (r *borrowsPostgresRepository) GetBorrowsByBookID(bookID uuid.UUID) ([]*model.Borrow, error) {
	var borrows []*model.Borrow
	if result := postgres.GetDB().
		Model(&model.Borrow{}).
		Joins("Book").
		Where("book_id = ?", bookID).
		Find(&borrows); result.Error != nil {
		return nil, fmt.Errorf("error retrieving borrows from database: %w", result.Error)
	}
	return borrows, nil
}

func (r *borrowsPostgresRepository) GetBorrowsNotBroughtByBookID(bookID uuid.UUID) ([]*model.Borrow, error) {
	var borrows []*model.Borrow
	if result := postgres.GetDB().
		Model(&model.Borrow{}).
		Joins("Book").
		Where("book_id = ?", bookID).
		Where("brought_date is null").
		Find(&borrows); result.Error != nil {
		return nil, fmt.Errorf("error retrieving borrows from database: %w", result.Error)
	}
	return borrows, nil
}
