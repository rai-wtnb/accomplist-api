package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/stretchr/testify/assert"
)

var researchCtrl = ResearchController{}

func TestResearchIndexUser(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/research", researchCtrl.ResearchIndex)

	url := fmt.Sprintf("/research?target=%v&req=%v",
		"user",
		userA.Description,
	)
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var results []models.User
	json.Unmarshal(respBodyByte, &results)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, userA.Description, results[0].Description, "invalid res data: Description")
	assert.Equal(t, userA.Name, results[0].Name, "invalid res data: Name")
}

func TestResearchIndexFeedback(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/research", researchCtrl.ResearchIndex)

	url := fmt.Sprintf("/research?target=%v&req=%v",
		"feedback",
		feedbackC.Body,
	)
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var results []models.Feedback
	json.Unmarshal(respBodyByte, &results)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, feedbackC.Title, results[0].Title, "invalid res data: Title")
	assert.Equal(t, feedbackC.Body, results[0].Body, "invalid res data: Body")
}
