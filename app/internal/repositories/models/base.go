package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel definition same as gorm.Model, but including other common columns
type BaseModel struct {
	gorm.Model
	ID         uint64    `gorm:"primaryKey;autoIncrement:true;column:id"`
	Identifier string    `gorm:"primaryKey;autoIncrement:false;column:identifier"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at"`
}
