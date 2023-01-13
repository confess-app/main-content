package model

import "gorm.io/gorm"

type Confession struct {
	gorm.Model
	ConfessionID string `json:"confession_id"`
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	Anonymous    bool   `json:"anonymous"`
	Content      string `json:"content" gorm:"type:text"`
	Category     string `json:"category"`
}

// TableName overrides the table name used by User to `user`
func (Confession) TableName() string {
	return "confession"
}
