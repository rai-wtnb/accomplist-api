package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/rai-wtnb/accomplist-api/models/repository"
)

type RelationController struct{}

// FollowIndex : GET /relations/follows/:id
func (RelationController) FollowIndex(c *gin.Context) {
	var r repository.RelationRepository
	var err error
	var followsAndFollowers models.FollowsAndFollowers
	id := c.Params.ByName("id")

	followids, err := r.GetFollowID(id)
	follows, err := r.GetRelationUser(followids)
	followerids, err := r.GetFollowerID(id)
	followers, err := r.GetRelationUser(followerids)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	followsAndFollowers.Follows = follows
	followsAndFollowers.Followers = followers

	c.JSON(200, followsAndFollowers)
}

// FollowerIndex : GET /relations/followers/:id
func (RelationController) FollowerIndex(c *gin.Context) {
	var r repository.RelationRepository
	var err error
	id := c.Params.ByName("id")

	ids, err := r.GetFollowerID(id)
	users, err := r.GetRelationUser(ids)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

// Count : GET /relations/count/:id
func (RelationController) Count(c *gin.Context) {
	var r repository.RelationRepository
	var err error
	id := c.Params.ByName("id")

	follows, err := r.GetFollowID(id)
	followers, err := r.GetFollowerID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := models.Count{}
	resp.FollowCount = len(follows)
	resp.FollowerCount = len(followers)

	c.JSON(200, resp)
}

// IsFollow : POST /relations/isfollow
func (RelationController) IsFollow(c *gin.Context) {
	var r repository.RelationRepository

	res, err := r.ConfirmRel(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}

// Create : POST /relations
func (RelationController) Create(c *gin.Context) {
	var r repository.RelationRepository

	err := r.CreateRelation(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "created the relation"})
}

// Delete : DELETE /relations/
func (RelationController) Delete(c *gin.Context) {
	var r repository.RelationRepository

	if err := r.DeleteRelation(c); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{"success": "deleted the relation"})
}
