package models

type User struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	Twitter     string `json:"twitter"`
	Description string `json:"description"`
	ImgPath     string `json:"img_path"`
	Lists       []List `json:"lists" binding:"dive"`
}
