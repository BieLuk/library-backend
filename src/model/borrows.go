package model

import (
	"github.com/google/uuid"
	"time"
)

type Borrow struct {
	DBEntity    `bson:",inline"`
	BookID      uuid.UUID  `bson:"book_id,omitempty"`
	Book        Book       `bson:"book,omitempty"`
	TakenDate   time.Time  `bson:"taken_date,omitempty"`
	BroughtDate *time.Time `bson:"brought_date,omitempty"`
}
