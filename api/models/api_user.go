package models

type ApiUser struct {
	ID          string `json:"id" binding:"required,max=20" gorm:"unique;primary_key;autoIncrement:false"`
	Name        string `json:"name" binding:"required,max=30"`
	Twitter     string `json:"twitter" binding:"max=20"`
	Description string `json:"description" binding:"max=200"`
	Img         string `json:"img"`
}
