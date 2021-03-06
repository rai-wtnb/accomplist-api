package repository

import (
	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
)

type ListRepository struct{}

type List models.List

// GetAll returns all Lists. used in controllers.Index()
func (ListRepository) GetAll() ([]models.List, error) {
	db := db.GetDB()
	var lists []models.List
	err := db.Table("lists").Scan(&lists).Error
	if err != nil {
		return nil, err
	}
	return lists, nil
}

// CreateList creates User model. used in controllers.Create()
func (ListRepository) CreateList(c *gin.Context) (List, error) {
	db := db.GetDB()
	var list List
	var err error

	err = c.BindJSON(&list)
	if err != nil {
		return list, err
	}
	err = db.Create(
		&List{
			UserID:  list.UserID,
			Content: list.Content,
			Done:    list.Done,
		}).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

// GetByUserID gets lists matches with user ID. used in controllers.IndexByUserID()
func (ListRepository) GetByUserID(id string) ([]models.List, error) {
	db := db.GetDB()
	var lists []models.List
	if err := db.Where("user_id = ?", id).Find(&lists).Error; err != nil {
		return lists, err
	}
	return lists, nil
}

// GetByListID get a list. used in contorollers.Show()
func (ListRepository) GetByListID(id string, c *gin.Context) (models.List, error) {
	db := db.GetDB()
	var list models.List
	if err := db.Where("ID = ?", id).Find(&list).Error; err != nil {
		return list, err
	}
	return list, nil
}

// UpdateByID updates a List. used in controllers.Update()
func (ListRepository) UpdateByID(id string, c *gin.Context) (models.List, error) {
	db := db.GetDB()
	var list models.List
	if err := db.Where("id = ?", id).First(&list).Error; err != nil {
		return list, err
	}
	if err := c.BindJSON(&list); err != nil {
		return list, err
	}
	db.Save(&list)
	return list, nil
}

// DeleteByID deletes a List matches with ID. used in controllers.Delete()
func (ListRepository) DeleteByID(id string) error {
	db := db.GetDB()
	var list List
	err := db.Where("id = ?", id).Delete(&list).Error
	if err != nil {
		return err
	}
	return nil
}
