package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/rai-wtnb/accomplist-api/models/repository"
	"github.com/rai-wtnb/accomplist-api/utils/mysession"
	"github.com/stretchr/testify/assert"
)

var feedbackCtrl = FeedbackController{}

func TestFeedbackIndex(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/feedbacks", feedbackCtrl.Index)

	req, _ := http.NewRequest("GET", "/feedbacks", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var feedbacks []models.Feedback
	json.Unmarshal(respBodyByte, &feedbacks)
	var resultA, resultB models.Feedback
	for _, feedback := range feedbacks {
		if feedback.Title == feedbackA.Title {
			resultA = feedback
		}
		if feedback.Title == feedbackB.Title {
			resultB = feedback
		}
	}

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, feedbackA.Body, resultA.Body, "invalid res data: Body of A")
	assert.Equal(t, feedbackB.ImgPath, resultB.ImgPath, "invalid res data: ImgPath of B")
}

func TestFeedbackCreate(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.POST("/feedbacks", feedbackCtrl.Create)

	// login
	var u repository.UserRepository
	sessionID := mysession.NewSessionID()
	u.SaveSession(userA.ID, sessionID)

	postC := fmt.Sprintf(`{"user_id":"%v","Img":"%v","Title":"%v","Body":"%v","Sess":"%v"}`,
		feedbackC.UserID,
		feedbackC.ImgPath,
		feedbackC.Title,
		feedbackC.Body,
		sessionID,
	)
	reqBody := strings.NewReader(postC)
	req, _ := http.NewRequest("POST", "/feedbacks", reqBody)
	r.ServeHTTP(w, req)

	testDb := db.Db
	var dbC models.Feedback
	testDb.Where("title = ?", feedbackC.Title).First(&dbC)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, feedbackC.Title, dbC.Title, "invalid db data: Title")
	assert.Equal(t, feedbackC.UserID, dbC.UserID, "invalid db data: UserID")
}

func TestFeedbackShow(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/feedbacks/:id", feedbackCtrl.Show)

	url := fmt.Sprintf("/feedbacks/%v", feedbackA.ListID)
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var resultA models.Feedback
	json.Unmarshal(respBodyByte, &resultA)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, feedbackA.ListID, resultA.ListID, "invalid res data: ListID")
	assert.Equal(t, feedbackA.Title, resultA.Title, "invalid res data: Title")
}

func TestFeedbackUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.PUT("/feedbacks/:id", feedbackCtrl.Update)

	// login
	var u repository.UserRepository
	sessionID := mysession.NewSessionID()
	u.SaveSession(userA.ID, sessionID)

	updateTitle := "title_a_update"
	updateBody := "body_a_update"

	putA := fmt.Sprintf(`{"user_id":"%v","Title":"%v","Body":"%v","Sess":"%v"}`,
		feedbackA.UserID,
		updateTitle,
		updateBody,
		sessionID,
	)
	reqBody := strings.NewReader(putA)
	url := fmt.Sprintf("/feedbacks/%v", feedbackA.ID)
	req, _ := http.NewRequest("PUT", url, reqBody)
	r.ServeHTTP(w, req)

	testDb := db.Db
	var dbA models.Feedback
	testDb.Where("id = ?", feedbackA.ID).First(&dbA)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, updateTitle, dbA.Title, "failed to update: Title")
	assert.Equal(t, updateBody, dbA.Body, "failed to update: Body")
}

func TestFeedbackDelete(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.DELETE("/feedbacks/:list-id", feedbackCtrl.Delete)

	url := fmt.Sprintf("/feedbacks/%v", listA.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	r.ServeHTTP(w, req)

	testDb := db.Db
	var dbA models.Feedback

	testDb.Where("title = ?", feedbackA.Title).First(&dbA)

	assert.Equal(t, 204, w.Code, "invalid StatusCode")
}
