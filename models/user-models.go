package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
}
