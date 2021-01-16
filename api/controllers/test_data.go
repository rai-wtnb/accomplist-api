package controllers

import "github.com/rai-wtnb/accomplist-api/models"

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
}

var listB = models.List{
	UserID:  "id_a",
	Content: "list_b",
}

var listC = models.List{
	UserID:  "id_a",
	Content: "list_c",
}

// feedback of listA
var feedbackA = models.Feedback{
	UserID:  "id_a",
	ImgPath: "img_a",
	Title:   "feedback_a",
	Body:    "feedback_a",
}

// feedback of listB
var feedbackB = models.Feedback{
	UserID:  "id_a",
	ImgPath: "img_b",
	Title:   "feedback_b",
	Body:    "feedback_b",
}
