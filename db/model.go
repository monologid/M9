package db

import (
	"time"

	"github.com/google/uuid"
)

// Model ...
type Model struct {
	ID        uuid.UUID `gorm:"column:_id;type:uuid;not null;unique;primaryKey;index;"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// GenerateID ...
func (m *Model) GenerateID() {
	m.ID, _ = uuid.NewRandom()
}
