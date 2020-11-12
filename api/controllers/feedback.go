package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rai-wtnb/accomplist-api/models/repository"
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
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, r)
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
