package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email    string    `gorm:"unique"`
	Username string    `gorm:"unique"`
	Password string
}
