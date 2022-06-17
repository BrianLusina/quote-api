package entities

import "time"

// BaseEntity is a base model for all entities
type BaseEntity struct {
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewBaseEntity returns a new base entity
func NewBaseEntity() BaseEntity {
	now := time.Now()
	return BaseEntity{
		CreatedAt: now.UTC().Round(time.Microsecond),
		UpdatedAt: now.UTC().Round(time.Microsecond),
	}
}
