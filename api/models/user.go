package models

type User struct {
	ID          string `json:"id" binding:"required,max=20" gorm:"unique;primary_key;autoIncrement:false"`
	Name        string `json:"name" binding:"required,max=30"`
	Email       string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password    string `json:"password" binding:"required,max=50"`
	Twitter     string `json:"twitter" binding:"max=20"`
	Description string `json:"description" binding:"max=200"`
	Img         string `json:"img"`
	Lists       []List `json:"-" binding:"dive"`
}
