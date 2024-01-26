package service

import "social-media/pkg/model"

type LikeService interface {
	CreateLike(like *model.Like) error
	GetLikeByID(likeID uint) (*model.Like, error)
	UpdateLike(like *model.Like) error
	DeleteLike(likeID uint) error
	GetAllLikes() ([]model.Like, error)
	GetLikesByPostID(postID uint) ([]model.Like, error)
	GetLikeCountByPostID(postID uint) (int, error)
}
