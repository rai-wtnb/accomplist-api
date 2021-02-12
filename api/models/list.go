package models

import "github.com/jinzhu/gorm"

type List struct {
	gorm.Model
	UserID   string   `gorm:"not null"  json:"user_id"`
	Content  string   `json:"content" binding:"required,max=100"`
	Done     bool     `json:"done" gorm:"dafault:false"`
	Feedback Feedback `json:"feedback" binding:"dive"`
	User     ApiUser  `json:"user"`
}
