package model

import (
	"github.com/google/uuid"
	"time"
)

type DBEntity struct {
	ID        *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()" bson:"_id,omitempty"`
	CreatedAt time.Time  `bson:"created_at,omitempty"`
	UpdatedAt time.Time  `bson:"updated_at,omitempty"`
}
