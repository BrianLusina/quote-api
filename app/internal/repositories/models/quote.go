package models

type Quote struct {
	BaseModel
	Quote  string `gorm:"column:quote;index:idx_quote"`
	Author string `gorm:"column:author;index:idx_quote"`
}

func (Quote) TableName() string {
	return "quotes"
}
