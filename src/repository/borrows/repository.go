package borrows

import (
	"github.com/BieLuk/library-backend/src/model"
	"github.com/google/uuid"
	"time"
)

type BorrowsRepository interface {
	CreateBorrow(book *model.Borrow) (*model.Borrow, error)
	GetBorrowsByBookID(bookID uuid.UUID) ([]*model.Borrow, error)
	GetBorrowsNotBroughtByBookID(bookID uuid.UUID) ([]*model.Borrow, error)
	UpdateBorrowBroughtDateByBookID(bookID uuid.UUID, broughtDate time.Time) error
}
