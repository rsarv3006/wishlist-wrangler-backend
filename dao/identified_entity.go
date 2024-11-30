package dao

import (
	"time"

	"github.com/google/uuid"
)

type IdentifiedEntity struct {
	ID        uuid.UUID  `gorm:"column:ID;primaryKey;autoIncrement"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (e *IdentifiedEntity) SetID(id uuid.UUID) {
	e.ID = id
}

func (e *IdentifiedEntity) GetID() uuid.UUID {
	return e.ID
}
