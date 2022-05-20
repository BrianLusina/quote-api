package models

import (
	"gorm.io/gorm"
)

// BaseModel definition same as gorm.Model, but including other common columns
type BaseModel struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey;autoIncrement:true;column:id"`
	Identifier string `gorm:"primaryKey;autoIncrement:false;column:identifier"`
}
