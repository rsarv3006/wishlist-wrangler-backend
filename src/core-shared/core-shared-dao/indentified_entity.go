package coreshareddao

import "time"

type IdentifiedEntity struct {
	ID        uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (e *IdentifiedEntity) SetID(id uint) {
	e.ID = id
}

func (e *IdentifiedEntity) GetID() uint {
	return e.ID
}
