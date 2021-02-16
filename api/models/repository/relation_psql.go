package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
)

type RelationRepository struct{}

type Relation models.Relation
type ConfirmRel models.ConfirmRel
type ApiUser models.ApiUser

// GetFollowID returns all IDs of users the user follows. used in controllers.FollwIndex()
func (RelationRepository) GetFollowID(id string) ([]string, error) {
	db := db.GetDB()
	var relations []Relation
	var ids []string

	if err := db.Where("follow_id = ?", id).Find(&relations).Error; err != nil {
		return ids, err
	}
	for _, relation := range relations {
		ids = append(ids, relation.FollowerID)
	}
	return ids, nil
}

// GetFollowerID returns all IDs of users follows the user. used in controllers.FollwerIndex()
func (RelationRepository) GetFollowerID(id string) ([]string, error) {
	db := db.GetDB()
	var relations []Relation
	var ids []string

	if err := db.Where("follower_id = ?", id).Find(&relations).Error; err != nil {
		return ids, err
	}
	for _, relation := range relations {
		ids = append(ids, relation.FollowID)
	}
	return ids, nil
}

// GetRelationUser returns all related Users. useded in controllers.FollowIndex()
func (RelationRepository) GetRelationUser(ids []string) ([]models.ApiUser, error) {
	db := db.GetDB()
	var apiUser models.ApiUser
	var apiUsers []models.ApiUser
	var err error

	for _, id := range ids {
		err = db.Table("users").Select("id, name, twitter, description, img").Where("id = ?", id).First(&apiUser).Error
		apiUsers = append(apiUsers, apiUser)
		apiUser = models.ApiUser{}
	}
	if err != nil {
		return nil, err
	}

	return apiUsers, err
}

// ConfirmRel confirm if relation exists. used in controllers.IsFollow()
func (RelationRepository) ConfirmRel(c *gin.Context) (ConfirmRel, error) {
	db := db.GetDB()
	var relation, result Relation
	var err error
	var confirmRel ConfirmRel

	err = c.BindJSON(&relation)
	if err != nil {
		return confirmRel, err
	}

	err = db.Where("follow_id = ? AND follower_id = ?",
		relation.FollowID,
		relation.FollowerID).Find(&result).Error
	if err != nil {
		confirmRel.IsFollow = false
	} else {
		confirmRel.IsFollow = true
	}

	return confirmRel, nil
}

// CreateRelation make new relations. used in controllers.Create()
func (RelationRepository) CreateRelation(c *gin.Context) error {
	db := db.GetDB()
	var relation Relation
	var err error

	err = c.BindJSON(&relation)
	if err != nil {
		return err
	}
	err = db.Create(&relation).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteRelation deletes the relation. used in controllers.Delete()
func (RelationRepository) DeleteRelation(c *gin.Context) error {
	db := db.GetDB()
	var relation Relation
	var err error

	err = c.BindJSON(&relation)
	if err != nil {
		return err
	}
	log.Println(relation)

	err = db.Where("follow_id = ? AND follower_id = ?",
		relation.FollowID,
		relation.FollowerID,
	).Delete(&relation).Error
	if err != nil {
		return err
	}
	return nil
}
