package repository

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(DB *gorm.DB) *CommentRepository {
	return &CommentRepository{DB: DB}
}

func (r *CommentRepository) CreateComment(comment *model.Comment) error {
	return r.DB.Create(comment).Error
}

func (r *CommentRepository) GetCommentByID(commentID uint) (*model.Comment, error) {
	var comment model.Comment
	err := r.DB.Preload("User").First(&comment, commentID).Error
	return &comment, err
}

func (r *CommentRepository) GetAllComments() ([]model.Comment, error) {
	var comments []model.Comment
	err := r.DB.Preload("User").Find(&comments).Error
	return comments, err
}

func (r *CommentRepository) UpdateComment(comment *model.Comment) error {
	return r.DB.Save(comment).Error
}

func (r *CommentRepository) GetCommentsByPost(postID uint) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.DB.Preload("User").Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}
func (r *CommentRepository) DeleteComment(commentID uint) error {
	return r.DB.Delete(&model.Comment{}, commentID).Error
}
