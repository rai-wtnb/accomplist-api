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
	"github.com/stretchr/testify/assert"
)

// TestMain is located in user_test.go

var listCtrl = ListController{}

func TestListIndex(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/lists", listCtrl.Index)

	req, _ := http.NewRequest("GET", "/lists", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var lists []models.List
	json.Unmarshal(respBodyByte, &lists)

	var resultA, resultB models.List
	for _, list := range lists {
		if list.Content == listA.Content {
			resultA = list
		}
		if list.Content == listB.Content {
			resultB = list
		}
	}

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, listA.Content, resultA.Content, "invalid res data: Content of A")
	assert.Equal(t, listA.UserID, resultA.UserID, "invalid res data: UserID of A")
	assert.Equal(t, listB.Done, resultB.Done, "invalid res data: Done of B")
	assert.Equal(t, listB.UserID, resultB.UserID, "invalid res data: UserID of B")

}

func TestListCreate(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.POST("/lists", listCtrl.Create)

	postC := fmt.Sprintf(`{"user_id":"%v","Content":"%v"}`,
		listC.UserID,
		listC.Content,
	)
	reqBody := strings.NewReader(postC)
	req, _ := http.NewRequest("POST", "/lists", reqBody)
	r.ServeHTTP(w, req)

	testDb := db.Db
	var resultC models.List
	testDb.Where("content = ?", listC.Content).First(&resultC)

	assert.Equal(t, 201, w.Code, "invalid StatusCode")
	assert.Equal(t, listC.Content, resultC.Content, "invalid db data: Content")
	assert.Equal(t, listC.UserID, resultC.UserID, "invalid db data: Content")
}

func TestListIndexByUserID(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/lists/specific/:id", listCtrl.IndexByUserID)

	req, _ := http.NewRequest("GET", "/lists/specific/id_a", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var lists []models.List
	json.Unmarshal(respBodyByte, &lists)

	var resultA, resultB models.List
	for _, list := range lists {
		if list.Content == listA.Content {
			resultA = list
		}
		if list.Content == listB.Content {
			resultB = list
		}
	}

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, listA.Content, resultA.Content, "invalid res data: Content of A")
	assert.Equal(t, listB.Content, resultB.Content, "invalid res data: Content of B")

}

func TestListShow(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/lists/specific/:id", listCtrl.Show)

	url := fmt.Sprintf("/lists/specific/%v", listA.ID)
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var resultA models.List
	json.Unmarshal(respBodyByte, &resultA)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, listA.ID, resultA.ID, "invalid res data")
}

func TestListUpdate(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/lists/specific/:id", listCtrl.Update)

	url := fmt.Sprintf("/lists/specific/%v", listA.ID)
	updateContent := "content_a_update"
	updateDone := false
	putA := fmt.Sprintf(`{"UserID":"%v","Content":"%v","Done":%t}`,
		listA.UserID,
		updateContent,
		updateDone,
	)
	reqBody := strings.NewReader(putA)
	req, _ := http.NewRequest("GET", url, reqBody)
	r.ServeHTTP(w, req)

	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var resultA models.List
	json.Unmarshal(respBodyByte, &resultA)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, updateContent, resultA.Content, "invalid db data: Content")
	assert.Equal(t, updateDone, resultA.Done, "invalid db data: Done")
}

func TestListDelete(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.DELETE("/lists/specific/:id", listCtrl.Delete)

	url := fmt.Sprintf("/lists/specific/%v", listA.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	r.ServeHTTP(w, req)

	resultDb := db.Db
	var resultA models.List
	resultDb.Where("id = ?", listA.ID).First(&resultA)

	assert.Equal(t, 204, w.Code, "invalid StatusCode")
	assert.Empty(t, resultA.Content, "failed to delete")

}
