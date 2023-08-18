package model

import (
	"github.com/google/uuid"
	"time"
)

type DBEntity struct {
	ID        *uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()" bson:"_id"`
	CreatedAt time.Time  `bson:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at"`
}
