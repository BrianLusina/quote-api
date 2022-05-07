package models

type Quote struct {
	BaseModel
	Quote  string `gorm:"column:quote;uniqueIndex:idx_quote"`
	Author string `gorm:"column:author;index:idx_author"`
}
