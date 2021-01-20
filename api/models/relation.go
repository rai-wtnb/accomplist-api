package models

type Relation struct {
	FollowID   string `gorm:"not null" json:"follow_id"`
	FollowerID string `gorm:"not null" json:"follower_id"`
}

type Count struct {
	FollowCount   int `json:"followCount"`
	FollowerCount int `json:"followerCount"`
}

type ConfirmRel struct {
	IsFollow bool `json:"isFollow"`
}
