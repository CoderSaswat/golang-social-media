package service_impl

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
)
import "social-media/pkg/repository"

type UserServiceImpl struct {
	userRepository *repository.UserRepository
}

func (s *UserServiceImpl) GetAllUsers() ([]model.User, error) {
	//TODO implement me
	return s.userRepository.GetAllUsers()
}

func NewUserServiceImpl(db *gorm.DB) *UserServiceImpl {
	userRepository := repository.NewUserRepository(db)
	return &UserServiceImpl{userRepository: userRepository}
}
func (s *UserServiceImpl) CreateUser(user *model.User) error {
	return s.userRepository.CreateUser(user)

}
func (s *UserServiceImpl) GetUserByID(userID uint) (*model.User, error) {
	return s.userRepository.GetUserByID(userID)
}
func (s *UserServiceImpl) UpdateUser(userID uint, updatedUser *model.User) error {
	return s.userRepository.UpdateUser(userID, updatedUser)
}
func (s *UserServiceImpl) DeleteUser(userID uint) error {
	return s.userRepository.DeleteUser(userID)
}
