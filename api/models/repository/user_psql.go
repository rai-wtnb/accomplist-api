package repository

import (
	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/crypto"
	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
)

type UserRepository struct{}

type User models.User

// GetAll gets all User. used in contorollers.Index()
func (UserRepository) GetAll() ([]models.ApiUser, error) {
	db := db.GetDB()
	var users []models.ApiUser
	if err := db.Table("users").Select("id, name, twitter, description, img").Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser creates User model. used in contorollers.Signup()
func (UserRepository) CreateUser(c *gin.Context) (User, error) {
	db := db.GetDB()
	var user User
	if err := c.BindJSON(&user); err != nil {
		return user, err
	}
	encryptedPassword := crypto.PasswordEncrypt(user.Password)
	if err := db.Create(
		&User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: encryptedPassword,
		}).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByEmail is used in controllers.Login()
func (UserRepository) GetByEmail(email string) User {
	db := db.GetDB()
	var user User
	db.First(&user, "email = ?", email)
	return user
}

// GetByID gets a User by ID. used in contorollers.Show()
func (UserRepository) GetByID(id string) (models.User, error) {
	db := db.GetDB()
	var me models.User
	if err := db.Where("id = ?", id).First(&me).Error; err != nil {
		return me, err
	}
	var lists []models.List
	db.Where("id = ?", id).First(&me)
	db.Model(&me).Related(&lists)
	me.Lists = lists

	return me, nil
}

// UpdateByID updates a User. used in contorollers.Update()
func (UserRepository) UpdateByID(id string, c *gin.Context) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	if err := c.BindJSON(&user); err != nil {
		return user, err
	}
	user.ID = string(id)
	db.Save(&user)
	return user, nil
}

// DeleteByID deletes a User matches with ID. used in contorollers.Delete()
func (UserRepository) DeleteByID(id string) error {
	db := db.GetDB()
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
