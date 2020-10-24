package db

import (
	"time"

	"github.com/google/uuid"
)

// Model is the default schema for database module
// This should be include in all schema model
type Model struct {
	ID        uuid.UUID `gorm:"column:_id;type:uuid;not null;unique;primaryKey;index;"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

// GenerateID will generate a random UUID and set it into the ID
func (m *Model) GenerateID() {
	m.ID, _ = uuid.NewRandom()
}
