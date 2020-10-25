package models

type Like struct {
	User     User     `json:"-" binding:"required" gorm:"unique"`
	List     List     `json:"-" binding:"required" gorm:"unique"`
}
