package dto

type FollowInput struct {
	UserId       uint   `json:"userId"`
	TargetUserId uint   `json:"targetUserId"`
	FollowType   string `json:"followType"`
}
