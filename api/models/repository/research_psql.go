package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/db"
)

type ResearchRepository struct{}

// GetResearchResult returns result of research. used in controllers.ResearchIndex()
func (ResearchRepository) GetResearchResult(c *gin.Context) (interface{}, error) {
	db := db.GetDB()
	var err error
	var users []ApiUser
	var feedbacks []Feedback

	target := c.Query("target")
	req := "%" + c.Query("req") + "%"

	if target == "user" {
		err = db.Table("users").Where("id LIKE ? OR name LIKE ? OR description LIKE ?",
			req, req, req).Find(&users).Error
	} else if target == "feedback" {
		err = db.Where("title LIKE ? OR body LIKE ?",
			req, req).Find(&feedbacks).Error
	} else {
		err = fmt.Errorf("Research Error: target is wrong")
	}

	if err != nil {
		return nil, err
	}

	if target == "user" {
		return users, nil
	}

	return feedbacks, nil
}
