package models

type Notification struct {
	Checked bool `json:"checked" binding:"required" gorm:"dafault:false"`
}
