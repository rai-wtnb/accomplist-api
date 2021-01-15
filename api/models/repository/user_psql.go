package repository

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/rai-wtnb/accomplist-api/utils/crypto"
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
		log.Println(err)
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
func (UserRepository) GetByEmail(email string) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Table("users").Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByID gets a User by ID. used in contorollers.Show()
func (UserRepository) GetByID(id string) (models.ApiUser, error) {
	db := db.GetDB()
	var me models.ApiUser
	if err := db.Table("users").Where("id = ?", id).Select("id, name, twitter, description, img").First(&me).Error; err != nil {
		return me, err
	}
	return me, nil
}

// UpdateByID updates a User. used in contorollers.Update()
func (UserRepository) UpdateByID(id string, userAndSession models.UserAndSession) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	user.Name = userAndSession.Name
	user.Twitter = userAndSession.Twitter
	user.Description = userAndSession.Description
	db.Save(&user)
	return user, nil
}

// SaveUrlByID saves url of image uploaded to s3. used in controllers.UpdateImage()
func (UserRepository) SaveUrlByID(id, url string) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	user.Img = url
	db.Save(&user)
	return user, nil
}

// DeleteByID deletes a User matches with ID. used in contorollers.Delete()
func (UserRepository) DeleteByID(id string) error {
	db := db.GetDB()
	var user models.User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// SaveSession save sessionID when user login or signup. used in contorollers.Login(), SignUp()
func (UserRepository) SaveSession(id, sessionID string) (models.User, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	user.SessionID = sessionID
	db.Save(&user)
	return user, nil
}

// getSession gets sessionID by userID.
func (UserRepository) GetSession(id string) (string, error) {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return user.SessionID, err
	}
	return user.SessionID, nil
}

// DeleteSession deletes sessionID when user logout. used in controllers.Logout()
func (UserRepository) DeleteSession(id string) error {
	db := db.GetDB()
	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	user.SessionID = ""
	db.Save(&user)
	return nil
}
