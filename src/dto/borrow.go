package dto

import (
	"github.com/google/uuid"
	"time"
)

type IsBookBorrowedResponse struct {
	IsBorrowed bool `json:"isBorrowed"`
}

type CreateBorrowRequest struct {
	BookID    uuid.UUID `json:"bookId" validate:"required"`
	TakenDate time.Time `json:"takenDate" validate:"required"`
}

type CreateBorrowResponse struct {
	ID uuid.UUID `json:"id"`
}

type ReturnBorrowRequest struct {
	BookID uuid.UUID `json:"book_id" validate:"required"`
}
