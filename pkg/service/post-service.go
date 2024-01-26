package service

import "social-media/pkg/model"

type PostService interface {
	CreatePost(post *model.Post) error
	GetPostByID(postID uint) (*model.Post, error)
	GetAllPosts() ([]model.Post, error)
	UpdatePost(post *model.Post) error
	DeletePost(postID uint) error
	GetAllPostsByUserID(userId uint) ([]model.Post, error)
}
