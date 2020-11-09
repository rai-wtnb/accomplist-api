package repository

import (
	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
)

type ListRepository struct{}

type List models.List

func (ListRepository) GetAll() ([]models.List, error) {
	db := db.GetDB()
	var lists []models.List
	if err := db.Table("lists").Scan(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}
