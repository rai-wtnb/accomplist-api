package models

import "github.com/jinzhu/gorm"

type Feedback struct {
	gorm.Model
	UserID  string `gorm:"not null" json:"user_id"`
	ListID  uint   `json:"list_id" gorm:"not null"`
	ImgPath string `json:"img"`
	Title   string `json:"title"`
	Body    string `json:"body" binding:"max=500"`
}

type FeedbackAndSession struct {
	UserID    string `json:"user_id"`
	ListID    uint   `json:"list_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	SessionID string `json:"sess"`
}
