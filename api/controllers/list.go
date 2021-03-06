package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/rai-wtnb/accomplist-api/models/repository"
)

type ListController struct{}

// Index : GET /lists
func (ListController) Index(c *gin.Context) {
	var l repository.ListRepository

	r, err := l.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, r)
}

// Create : POST /lists/
func (ListController) Create(c *gin.Context) {
	var l repository.ListRepository

	r, err := l.CreateList(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, r)
}

// IndexByUserID : GET /lists/specific/:id
func (ListController) IndexByUserID(c *gin.Context) {
	var l repository.ListRepository
	id := c.Params.ByName("id")

	list, err := l.GetByUserID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, list)
}

// Show : GET /lists/specific/:id
func (ListController) Show(c *gin.Context) {
	var err error
	var listAndFeedback models.ListAndFeedback
	var l repository.ListRepository
	var u repository.UserRepository
	var f repository.FeedbackRepository

	id := c.Params.ByName("id")

	list, err := l.GetByListID(id, c)
	userID := list.UserID
	feedback, err := f.GetByListID(id)
	user, err := u.GetByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	listAndFeedback.ID = list.ID
	listAndFeedback.UserID = list.UserID
	listAndFeedback.Content = list.Content
	listAndFeedback.Done = list.Done
	listAndFeedback.Feedback = feedback
	listAndFeedback.User = user

	c.JSON(200, listAndFeedback)
}

// Update : PUT /lists/specific/:id
func (ListController) Update(c *gin.Context) {
	var l repository.ListRepository
	id := c.Params.ByName("id")

	r, err := l.UpdateByID(id, c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, r)
}

// Delete : DELETE /lists/specific/:id
func (ListController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var l repository.ListRepository

	if err := l.DeleteByID(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{"success": "deleted the list"})
}
