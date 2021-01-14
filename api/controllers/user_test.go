package controllers

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/rai-wtnb/accomplist-api/db"
	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/stretchr/testify/assert"
)

var userA = models.User{
	ID:          "id_a",
	Name:        "name_a",
	Email:       "email_a",
	Password:    "pass_a",
	Twitter:     "twitter_a",
	Description: "description_a",
	Img:         "img_a",
}

var userB = models.User{
	ID:          "id_b",
	Name:        "name_b",
	Email:       "email_b",
	Password:    "pass_b",
	Twitter:     "twitter_b",
	Description: "description_b",
	Img:         "img_b",
}

func getDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	var err error
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("postgres", db)
	if err != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}

func TestIndex(t *testing.T) {
	gdb, mock, err := getDBMock()
	if err != nil {
		t.Fatal(err)
	}
	defer gdb.Close()
	gdb.LogMode(true)

	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"twitter",
		"description",
		"img",
	}).
		AddRow(
			userA.ID,
			userA.Name,
			userA.Twitter,
			userA.Description,
			userA.Img,
		).
		AddRow(
			userB.ID,
			userB.Name,
			userB.Twitter,
			userB.Description,
			userB.Img,
		)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"`)).
		WillReturnRows(rows)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	db.Init()
	defer db.Close()
	ctrl := UserController{}
	r.GET("/users", ctrl.Index)
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code, "invalid StatusCode")
}
