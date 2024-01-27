package service_impl

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"social-media/pkg/dto"
	"social-media/pkg/model"
	"social-media/pkg/utils"
)
import "social-media/pkg/repository"

type UserServiceImpl struct {
	userRepository   *repository.UserRepository
	followRepository *repository.FollowRepository
}

func (s *UserServiceImpl) GetFollowingsAndFollowersInfo(userID uint) (dto.FollowingOutput, error) {
	isExists := s.UserExistsById(userID)
	if !isExists {
		return dto.FollowingOutput{}, errors.New(fmt.Sprintf("user with this id %v does not exists", userID))
	}
	followings := s.followRepository.GetFollowByUserId(userID)
	followingsUser := utils.Map(func(follow model.Follow) model.User {
		return follow.TargetUser
	}, followings)
	followers := s.followRepository.GetFollowByTargetUserId(userID)
	followersUsers := utils.Map(func(follow model.Follow) model.User {
		return follow.User
	}, followers)
	fmt.Println(followingsUser, followersUsers)
	followingOutput := dto.FollowingOutput{
		FollowingsCount: len(followingsUser),
		FollowersCount:  len(followersUsers),
		Followings:      followingsUser,
		Followers:       followersUsers,
	}
	return followingOutput, nil
}

func (s *UserServiceImpl) FollowOrUnfollow(input dto.FollowInput) error {

	//_, err := strconv.Atoi("g")
	//if err != nil {
	//	return err
	//}

	if !s.UserExistsById(input.UserId) {
		return errors.New(fmt.Sprintf("user with this id %v does not exits", input.UserId))
	}
	if !s.UserExistsById(input.TargetUserId) {
		return errors.New(fmt.Sprintf("target user with this id %v does not exits", input.TargetUserId))
	}
	if input.UserId == input.TargetUserId {
		return errors.New("can't follow or unfollow self")
	}
	isFollowing := s.followRepository.IsFollowing(input.UserId, input.TargetUserId)
	if input.FollowType == "follow" {
		if isFollowing {
			return errors.New("can't follow again, since you are already following")
		}

		follow := model.Follow{
			UserID:       input.UserId,
			TargetUserID: input.TargetUserId,
		}
		err := s.followRepository.Follow(&follow)
		if err != nil {
			return err
		}
	} else if input.FollowType == "unfollow" {
		if !isFollowing {
			return errors.New("can't unfollow , since you are not following to this person")
		}
		follow := s.followRepository.GetByUserIdAndTargetUserId(input.UserId, input.TargetUserId)
		err := s.followRepository.DeleteFollowById(follow.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *UserServiceImpl) UserExistsById(userID uint) bool {
	return s.userRepository.UserExistsById(userID)
}

func (s *UserServiceImpl) GetAllUsers() ([]model.User, error) {
	//TODO implement me
	return s.userRepository.GetAllUsers()
}

func NewUserServiceImpl(db *gorm.DB) *UserServiceImpl {
	userRepository := repository.NewUserRepository(db)
	followRepository := repository.NewFollowRepository(db)
	return &UserServiceImpl{
		userRepository:   userRepository,
		followRepository: followRepository,
	}
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
