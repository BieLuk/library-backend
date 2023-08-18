package model

import (
	"github.com/google/uuid"
	"time"
)

type Borrow struct {
	DBEntity    `bson:",inline"`
	BookID      uuid.UUID  `bson:"book_id"`
	Book        Book       `bson:",omitempty"`
	TakenDate   time.Time  `bson:"taken_date"`
	BroughtDate *time.Time `bson:"brought_date"`
}
