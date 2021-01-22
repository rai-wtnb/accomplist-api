package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/models/repository"
)

type ResearchController struct{}

func (ResearchController) ResearchIndex(c *gin.Context) {
	var r repository.ResearchRepository

	res, err := r.GetResearchResult(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, res)
}
