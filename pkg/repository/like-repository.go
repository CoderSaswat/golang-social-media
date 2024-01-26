package repository

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
)

type LikeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *LikeRepository {
	return &LikeRepository{DB: db}
}

func (r *LikeRepository) CreateLike(like *model.Like) error {
	return r.DB.Create(like).Error
}

func (r *LikeRepository) GetLikeByID(likeID uint) (*model.Like, error) {
	var like model.Like
	err := r.DB.Preload("User").First(&like, likeID).Error
	return &like, err
}

func (r *LikeRepository) UpdateLike(like *model.Like) error {
	return r.DB.Save(like).Error
}

func (r *LikeRepository) DeleteLike(likeID uint) error {
	return r.DB.Delete(&model.Like{}, likeID).Error
}

func (r *LikeRepository) GetAllLikes() ([]model.Like, error) {
	var likes []model.Like
	err := r.DB.Preload("User").Find(&likes).Error
	return likes, err
}

func (r *LikeRepository) GetLikesByPostID(postID uint) ([]model.Like, error) {
	var likes []model.Like
	err := r.DB.Preload("User").Where("post_id = ?", postID).Find(&likes).Error
	return likes, err
}

func (r *LikeRepository) GetLikeCountByPostID(postID uint) (int, error) {
	var count int64
	err := r.DB.Model(&model.Like{}).Where("post_id = ?", postID).Count(&count).Error
	return int(count), err
}
