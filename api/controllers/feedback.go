package controllers

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/models/repository"
	"github.com/rai-wtnb/accomplist-api/utils/s3"
)

type FeedbackController struct{}

// Index : GET /feedbacks
func (FeedbackController) Index(c *gin.Context) {
	var f repository.FeedbackRepository
	r, err := f.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, r)
	}
}

// Create : POST /feedbacks
func (FeedbackController) Create(c *gin.Context) {
	var f repository.FeedbackRepository
	r, err := f.CreateFeedback(c)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, r)
	}
}

// Show : GET /feedbacks/:id
func (FeedbackController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var f repository.FeedbackRepository
	feedback, err := f.GetByListID(id)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, feedback)
	}
}

// Update : PUT /feedbacks/:id
func (FeedbackController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var f repository.FeedbackRepository
	r, err := f.UpdateByID(id, c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, r)
	}
}

// UpdateImg : PUT /feedbacks/:list-id/img
func (FeedbackController) UpdateImgByListID(c *gin.Context) {
	id := c.Params.ByName("id")
	var f repository.FeedbackRepository
	img, header, err := c.Request.FormFile("img")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

		// upload to s3
		buf := bytes.NewBuffer(nil)
		_, err := io.Copy(buf, img)
		url, err := s3.Upload(buf.Bytes(), header.Filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		// save url in db
		_, err = f.SaveUrlByListID(id, url)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(200, err)
	}
}

//Delete: DELETE /feedbacks/:id
func (FeedbackController) Delete(c *gin.Context) {
	id := c.Params.ByName("list-id")
	var f repository.FeedbackRepository
	err := f.DeleteByListID(id)
	if err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "達成フィードバックを削除しました"})
	return
}
