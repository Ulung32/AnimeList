package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Anime struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title       string
	Synopsis    string
	RelaseDate  time.Time
	RiviewSum   int
	RiviewCount int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
