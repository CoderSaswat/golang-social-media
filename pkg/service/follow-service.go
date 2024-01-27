package service

type FollowService interface {
	isFollowing(userID uint, followerID uint) bool
}
