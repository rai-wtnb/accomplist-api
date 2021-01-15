package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/rai-wtnb/accomplist-api/utils/crypto"
	"github.com/stretchr/testify/assert"
)

var userA = models.User{
	ID:          "id_a",
	Name:        "name_a",
	Email:       "email@aaa.com",
	Password:    "pass_a",
	Twitter:     "twitter_a",
	Description: "description_a",
	Img:         "img_a",
}

var userB = models.User{
	ID:          "id_b",
	Name:        "name_b",
	Email:       "email@bbb.com",
	Password:    "pass_b",
	Twitter:     "twitter_b",
	Description: "description_b",
	Img:         "img_b",
}

var userC = models.User{
	ID:          "id_c",
	Name:        "name_c",
	Email:       "email@ccc.com",
	Password:    "pass_c",
	Twitter:     "twitter_c",
	Description: "description_c",
	Img:         "img_c",
}

var ctrl = UserController{}

func TestMain(m *testing.M) {
	// before all
	gin.SetMode(gin.TestMode)
	db.Init()
	defer db.Close()

	testDb := db.Db
	userA.Password = crypto.PasswordEncrypt(userA.Password)
	userB.Password = crypto.PasswordEncrypt(userB.Password)
	testDb.Create(&userA)
	testDb.Create(&userB)

	code := m.Run()

	// after all
	testDb.Exec("DELETE FROM users WHERE id = ? OR id = ? OR id = ?", userA.ID, userB.ID, userC.ID)
	os.Exit(code)
}

func TestSignup(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.POST("/signup", ctrl.Signup)

	postC := fmt.Sprintf(`{"ID":"%v","Name":"%v","Email":"%v","Password":"%v"}`,
		userC.ID,
		userC.Name,
		userC.Email,
		userC.Password)
	reqBody := strings.NewReader(postC)
	req, _ := http.NewRequest("POST", "/signup", reqBody)
	r.ServeHTTP(w, req)

	testDb := db.Db
	var resultC models.User
	testDb.Where("id = ?", "id_c").First(&resultC)

	assert.Equal(t, 201, w.Code, "invalid StatusCode")
	assert.Equal(t, userC.ID, resultC.ID, "invalid DB data: ID")
	if err := crypto.Verify(resultC.Password, userC.Password); err != nil {
		t.Fatal(err)
	}
}

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.POST("/login", ctrl.Login)

	loginA := fmt.Sprintf("email=%v;password=%v", userA.Email, "pass_a")
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(loginA))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;")
	r.ServeHTTP(w, req)
	resp := w.Result()

	testDb := db.Db
	var dbA models.User
	testDb.Where("id = ?", "id_a").First(&dbA)

	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	result := gin.H{}
	json.Unmarshal(respBodyByte, &result)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, dbA.SessionID, result["sessionID"], "invalid sessionID")
	assert.Equal(t, dbA.ID, result["userID"], "invalid userID")
}

func TestIndex(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	r.GET("/users", ctrl.Index)
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)
	resp := w.Result()
	respBodyByte, _ := ioutil.ReadAll(resp.Body)
	var users []models.User
	json.Unmarshal(respBodyByte, &users)

	var resultA, resultB models.User
	for _, user := range users {
		if user.ID == "id_a" {
			resultA = user
		}
		if user.ID == "id_b" {
			resultB = user
		}
	}

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
	assert.Equal(t, userA.ID, resultA.ID, "invalid response data: ID of A")
	assert.Equal(t, userB.ID, resultB.ID, "invalid response data: ID of B")
	assert.Equal(t, userA.Img, resultA.Img, "invalid response data: Img of A")
	assert.Equal(t, userB.Description, resultB.Description, "invalid response data: Description of B")
}
