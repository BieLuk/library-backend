package borrows

import (
	"github.com/BieLuk/library-backend/src/model"
	"github.com/google/uuid"
)

type BorrowsRepository interface {
	CreateBorrow(book *model.Borrow) (*model.Borrow, error)
	GetBorrowByBookID(bookID uuid.UUID) (*model.Borrow, error)
}
