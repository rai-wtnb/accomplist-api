package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/rai-wtnb/accomplist-api/models/repository"
	"github.com/rai-wtnb/accomplist-api/utils/s3"
)

type FeedbackController struct{}

// Index : GET /feedbacks
func (FeedbackController) Index(c *gin.Context) {
	var f repository.FeedbackRepository

	r, err := f.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, r)
}

// Create : POST /feedbacks
func (FeedbackController) Create(c *gin.Context) {
	var feedbackAndSession models.FeedbackAndSession
	var err error

	if err = c.BindJSON(&feedbackAndSession); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var f repository.FeedbackRepository
	var u repository.UserRepository
	dbSessionID, _ := u.GetSession(feedbackAndSession.UserID)

	// validation
	if feedbackAndSession.SessionID == "" || dbSessionID != feedbackAndSession.SessionID {
		log.Println("wrong sessionID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong sessionID"})
		return
	}

	r, err := f.CreateFeedback(feedbackAndSession)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, r)
}

// Show : GET /feedbacks/:id
func (FeedbackController) Show(c *gin.Context) {
	var f repository.FeedbackRepository
	id := c.Params.ByName("id")

	feedback, err := f.GetByListID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, feedback)
}

// Update : PUT /feedbacks/:id
func (FeedbackController) Update(c *gin.Context) {
	var feedbackAndSession models.FeedbackAndSession
	var err error

	err = c.BindJSON(&feedbackAndSession)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Params.ByName("id")
	var f repository.FeedbackRepository
	var u repository.UserRepository
	dbSessionID, _ := u.GetSession(feedbackAndSession.UserID)

	// validation
	if feedbackAndSession.SessionID == "" || dbSessionID != feedbackAndSession.SessionID {
		log.Println("wrong sessionID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong sessionID"})
		return
	}

	r, err := f.UpdateByID(id, feedbackAndSession)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, r)
}

// UpdateImgByListID : PUT /feedbacks/:list-id/img
func (FeedbackController) UpdateImgByListID(c *gin.Context) {
	id := c.Params.ByName("id")
	var f repository.FeedbackRepository
	var err error

	img, header, err := c.Request.FormFile("img")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// upload to s3
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, img)
	url, err := s3.Upload(buf.Bytes(), header.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save url in db
	_, err = f.SaveUrlByListID(id, url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, err)
}

// Delete : DELETE /feedbacks/:id
func (FeedbackController) Delete(c *gin.Context) {
	id := c.Params.ByName("list-id")
	var f repository.FeedbackRepository

	err := f.DeleteByListID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{"success": "deleted the feedback"})
}
