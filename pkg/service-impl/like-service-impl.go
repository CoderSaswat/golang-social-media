package service_impl

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
	"social-media/pkg/repository"
)

type LikeServiceImpl struct {
	likeRepository *repository.LikeRepository
}

func NewLikeServiceImpl(db *gorm.DB) *LikeServiceImpl {
	likeRepository := repository.NewLikeRepository(db)
	return &LikeServiceImpl{likeRepository: likeRepository}
}

func (s *LikeServiceImpl) CreateLike(like *model.Like) error {
	return s.likeRepository.CreateLike(like)
}

func (s *LikeServiceImpl) GetLikeByID(likeID uint) (*model.Like, error) {
	return s.likeRepository.GetLikeByID(likeID)
}

func (s *LikeServiceImpl) UpdateLike(like *model.Like) error {
	return s.likeRepository.UpdateLike(like)
}

func (s *LikeServiceImpl) DeleteLike(likeID uint) error {
	return s.likeRepository.DeleteLike(likeID)
}

func (s *LikeServiceImpl) GetAllLikes() ([]model.Like, error) {
	return s.likeRepository.GetAllLikes()
}

func (s *LikeServiceImpl) GetLikesByPostID(postID uint) ([]model.Like, error) {
	return s.likeRepository.GetLikesByPostID(postID)
}

func (s *LikeServiceImpl) GetLikeCountByPostID(postID uint) (int, error) {
	return s.likeRepository.GetLikeCountByPostID(postID)
}
