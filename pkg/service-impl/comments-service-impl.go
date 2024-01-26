package service_impl

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
	"social-media/pkg/repository"
)

type CommentServiceImpl struct {
	commentRepository *repository.CommentRepository
}

func NewCommentServiceImpl(db *gorm.DB) *CommentServiceImpl {
	commentRepository := repository.NewCommentRepository(db)
	return &CommentServiceImpl{commentRepository: commentRepository}
}

func (s *CommentServiceImpl) CreateComment(comment *model.Comment) error {
	return s.commentRepository.CreateComment(comment)
}

func (s *CommentServiceImpl) GetCommentByID(commentID uint) (*model.Comment, error) {
	return s.commentRepository.GetCommentByID(commentID)
}

func (s *CommentServiceImpl) GetAllComments() ([]model.Comment, error) {
	return s.commentRepository.GetAllComments()
}

func (s *CommentServiceImpl) UpdateComment(comment *model.Comment) error {
	return s.commentRepository.UpdateComment(comment)
}

func (s *CommentServiceImpl) DeleteComment(commentID uint) error {
	return s.commentRepository.DeleteComment(commentID)
}

func (s *CommentServiceImpl) GetCommentsByPost(postID uint) ([]model.Comment, error) {
	return s.commentRepository.GetCommentsByPost(postID)
}
