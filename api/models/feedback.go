package models

import "github.com/jinzhu/gorm"

type Feedback struct {
	gorm.Model
	ListID  uint   `json:"list_id" gorm:"not null"`
	ImgPath string `json:"img"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}
