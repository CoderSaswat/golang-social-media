package service

import (
	"social-media/pkg/dto"
	"social-media/pkg/model"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(userID uint) (*model.User, error)
	UpdateUser(userID uint, updatedUser *model.User) error
	DeleteUser(userID uint) error
	GetAllUsers() ([]model.User, error)
	UserExistsById(userID uint) bool
	FollowOrUnfollow(input dto.FollowInput) error
	GetFollowingsAndFollowersInfo(userID uint) (dto.FollowingOutput, error)
}
