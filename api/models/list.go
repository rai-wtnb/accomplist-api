package models

import "github.com/jinzhu/gorm"

type List struct {
	gorm.Model
	Content  string   `json:"content" binding:"required"`
	User     User     `json:"-" binding:"required"`
	UserID   uint     `json:"user_id" binding:"required"`
	Done     bool     `json:done gorm:"dafault:false"`
	Feedback Feedback `json:"-"`
}

// 100文字・200文字制限
