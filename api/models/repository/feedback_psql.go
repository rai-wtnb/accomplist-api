package repository

import (
	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
)

type FeedbackRepository struct{}

type Feedback models.Feedback

// GetAll gets all Feedbacks. used in controllers.Index()
func (FeedbackRepository) GetAll() ([]models.Feedback, error) {
	db := db.GetDB()
	var feedbacks []models.Feedback
	err := db.Table("feedbacks").Scan(&feedbacks).Error
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

// CreateFeedback creates Feedback model. used in controllers.Create()
func (FeedbackRepository) CreateFeedback(c *gin.Context) (Feedback, error) {
	db := db.GetDB()
	var feedback Feedback
	err := c.BindJSON(&feedback)
	if err != nil {
		return feedback, err
	}
	if err := db.Create(
		&Feedback{
			UserID:  feedback.UserID,
			ListID:  feedback.ListID,
			ImgPath: feedback.ImgPath,
			Title:   feedback.Title,
			Body:    feedback.Body,
		}).Error; err != nil {
		return feedback, err
	}
	return feedback, nil
}

// GetByListID get a Feedback matched with list_ID. used in controllers.Show()
func (FeedbackRepository) GetByListID(id string) (models.Feedback, error) {
	db := db.GetDB()
	var feedback models.Feedback
	if err := db.Where("list_id = ?", id).Find(&feedback).Error; err != nil {
		return feedback, err
	}
	return feedback, nil
}

// UpdateByID updates Feedback. used in controllers.Update()
func (FeedbackRepository) UpdateByID(id string, c *gin.Context) (models.Feedback, error) {
	db := db.GetDB()
	var feedback models.Feedback
	if err := db.Where("id = ?", id).First(&feedback).Error; err != nil {
		return feedback, err
	}
	if err := c.BindJSON(&feedback); err != nil {
		return feedback, err
	}
	db.Save(&feedback)
	return feedback, nil
}

// DeleteByListID deletes a feedback batches with ID. used in controllers.Delete()
func (FeedbackRepository) DeleteByListID(id string) error {
	db := db.GetDB()
	var feedback Feedback
	err := db.Where("list_id = ?", id).Delete(&feedback).Error
	if err != nil {
		return err
	}
	return nil
}
