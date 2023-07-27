package model

import (
	"github.com/google/uuid"
	"time"
)

type Borrow struct {
	DBEntity
	BookID      uuid.UUID
	Book        Book
	TakenDate   time.Time
	BroughtDate *time.Time
}
