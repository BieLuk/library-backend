package dto

import (
	"github.com/google/uuid"
	"time"
)

type IsBookBorrowedResponse struct {
	IsBorrowed bool `json:"isBorrowed"`
}

type CreateBorrowRequest struct {
	BookID    uuid.UUID `json:"bookId"`
	TakenDate time.Time `json:"takenDate"`
}

type CreateBorrowResponse struct {
	ID uuid.UUID `json:"id"`
}
