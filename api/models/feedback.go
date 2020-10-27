package models

type Feedback struct {
	ID      uint   `json:"id" binding:"required" gorm:"primary_key"`
	ImgPath string `json:"img"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}
