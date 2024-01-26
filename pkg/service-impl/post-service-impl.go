package service_impl

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
	"social-media/pkg/repository"
)

type PostServiceImpl struct {
	postRepository *repository.PostRepository
}

func NewPostServiceImpl(db *gorm.DB) *PostServiceImpl {
	postRepository := repository.NewPostRepository(db)
	return &PostServiceImpl{postRepository: postRepository}
}
func (s *PostServiceImpl) GetAllPostsByUserID(userId uint) ([]model.Post, error) {
	return s.postRepository.GetAllPostsByUserID(userId)
}

func (s *PostServiceImpl) CreatePost(post *model.Post) error {
	return s.postRepository.CreatePost(post)
}

func (s *PostServiceImpl) GetPostByID(postID uint) (*model.Post, error) {
	return s.postRepository.GetPostByID(postID)
}

func (s *PostServiceImpl) GetAllPosts() ([]model.Post, error) {
	return s.postRepository.GetAllPosts()
}

func (s *PostServiceImpl) UpdatePost(post *model.Post) error {
	return s.postRepository.UpdatePost(post)
}

func (s *PostServiceImpl) DeletePost(postID uint) error {
	return s.postRepository.DeletePost(postID)
}
