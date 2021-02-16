package repository

import (
	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
)

type FeedbackRepository struct{}

type Feedback models.Feedback

// GetAll gets all Feedbacks. used in controllers.Index()
func (FeedbackRepository) GetAll() ([]models.FeedbackAndUser, error) {
	db := db.GetDB()
	var feedbacks []models.FeedbackAndUser
	err := db.Table("feedbacks").Scan(&feedbacks).Error
	if err != nil {
		return nil, err
	}
	return feedbacks, nil
}

// CreateFeedback creates Feedback model. used in controllers.Create()
func (FeedbackRepository) CreateFeedback(feedbackAndSession models.FeedbackAndSession) (Feedback, error) {
	db := db.GetDB()
	var feedback Feedback
	if err := db.Create(
		&Feedback{
			UserID: feedbackAndSession.UserID,
			ListID: feedbackAndSession.ListID,
			Title:  feedbackAndSession.Title,
			Body:   feedbackAndSession.Body,
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
func (FeedbackRepository) UpdateByID(id string, feedbackAndSession models.FeedbackAndSession) (models.Feedback, error) {
	db := db.GetDB()
	var feedback models.Feedback
	if err := db.Where("id = ?", id).First(&feedback).Error; err != nil {
		return feedback, err
	}

	feedback.UserID = feedbackAndSession.UserID
	feedback.ListID = feedbackAndSession.ListID
	feedback.Title = feedbackAndSession.Title
	feedback.Body = feedbackAndSession.Body
	db.Save(&feedback)
	return feedback, nil
}

// SaveUrlByID saves url of image uploaded to s3. used in controllers.UpdateImage()
func (FeedbackRepository) SaveUrlByListID(id, url string) (models.Feedback, error) {
	db := db.GetDB()
	var feedback models.Feedback
	if err := db.Where("list_id = ?", id).First(&feedback).Error; err != nil {
		return feedback, err
	}
	feedback.ImgPath = url
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
