package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Name         string
	IsVerified   bool
	IsAdmin bool     `gorm:"default:false" json:"is_admin"`
	OrgID        uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
