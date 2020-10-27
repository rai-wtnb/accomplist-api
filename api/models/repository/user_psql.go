package repository

import (
	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/crypto"
	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
)

type UserRepository struct{}

type User models.User

// GetAll is gets all User. used in contorollers.Index()
func (UserRepository) GetAll() ([]models.User, error) {
	db := db.GetDB()
	var users []models.User
	if err := db.Table("users").Select("name, id, email, password, twitter, description, img").Scan(&users).Error; err != nil {
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
	if err := db.Create(&User{Name: user.Name, Email: user.Email, Password: encryptedPassword}).Error; err != nil {
		return user, err
	}
	return user, nil
}

// getByEmail is used in controllers.Login()
func (UserRepository) GetByEmail(email string) User {
	db := db.GetDB()
	var user User
	db.First(&user, "email = ?", email)
	return user
}

// GetByID gets a User by ID. used in contorollers.Show()
func (UserRepository) GetByID(id int) (models.User, error) {
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
func (UserRepository) UpdateByID(id int, c *gin.Context) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	if err := c.BindJSON(&user); err != nil {
		return user, err
	}
	user.ID = uint(id)
	db.Save(&user)

	return user, nil
}

// DeleteByID deletes a User by ID. used in contorollers.Delete()
func (UserRepository) DeleteByID(id int) error {
	db := db.GetDB()
	var user User

	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
