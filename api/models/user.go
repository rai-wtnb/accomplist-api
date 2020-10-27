package models

type User struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password    string `json:"password" binding:"required"`
	Twitter     string `json:"twitter"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Lists       []List `json:"lists" binding:"dive"`
}
