package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/rai-wtnb/accomplist-api/models/repository"
)

type ListController struct{}

// Index : GET /lists
func (ListController) Index(c *gin.Context){
	var l repository.ListRepository
	r, err := l.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, r)
	}
}

// Create : POST /lists/
func (ListController) Create(c *gin.Context){
	var l repository.ListRepository
	r, err := l.CreateList(c)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, r)
	}
}

// Show : GET /lists/specific/:id
func (ListController) IndexByUserID(c *gin.Context) {
	id := c.Params.ByName("id")
	var l repository.ListRepository
	list, err := l.GetByUserID(id)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, list)
	}
}

// Show : GET /lists/specific/:id
func (ListController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var l repository.ListRepository
	r, err := l.GetByListID(id, c)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, r)
	}
}

// Update : PUT /lists/specific/:id
func (ListController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var l repository.ListRepository
	r, err := l.UpdateByID(id, c)
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, r)
	}
}

// Delete : DELETE /lists/specific/:id
func (ListController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var l repository.ListRepository
	err := l.DeleteByID(id)
	if err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "リストを削除しました"})
	return
}
