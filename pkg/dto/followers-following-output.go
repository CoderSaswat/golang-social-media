package dto

import "social-media/pkg/model"

type FollowingOutput struct {
	FollowingsCount int          `json:"followingsCount"`
	FollowersCount  int          `json:"followersCount"`
	Followings      []model.User `json:"followings"`
	Followers       []model.User `json:"followers"`
}
