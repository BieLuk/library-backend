package model

import (
	"github.com/google/uuid"
	"time"
)

type DBEntity struct {
	ID        *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
