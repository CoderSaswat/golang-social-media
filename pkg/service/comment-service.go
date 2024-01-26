package service

import "social-media/pkg/model"

type CommentService interface {
	CreateComment(comment *model.Comment) error
	GetCommentByID(commentID uint) (*model.Comment, error)
	GetAllComments() ([]model.Comment, error)
	UpdateComment(comment *model.Comment) error
	DeleteComment(commentID uint) error
	GetCommentsByPost(postID uint) ([]model.Comment, error)
}
