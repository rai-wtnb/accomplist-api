package controllers

import (
	"github.com/rai-wtnb/accomplist-api/models"
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

var listA = models.List{
	UserID:  "id_a",
	Content: "list_a",
	Done:    true,
}

var listB = models.List{
	UserID:  "id_a",
	Content: "list_b",
	Done:    true,
}

var listC = models.List{
	UserID:  "id_a",
	Content: "list_c",
}

// feedback of listA
var feedbackA = models.Feedback{
	UserID:  "id_a",
	ImgPath: "img_a",
	Title:   "title_a",
	Body:    "body_a",
}

// feedback of listB
var feedbackB = models.Feedback{
	UserID:  "id_a",
	ImgPath: "img_b",
	Title:   "title_b",
	Body:    "body_b",
}

// feedback of listC
var feedbackC = models.Feedback{
	UserID:  "id_a",
	ImgPath: "img_c",
	Title:   "title_c",
	Body:    "body_c",
}
