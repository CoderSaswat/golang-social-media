package repository

// repositories/user_repository.go

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	err := r.DB.Preload("Address").First(&user, userID).Error
	return &user, err
}

func (r *UserRepository) UpdateUser(userID uint, updatedUser *model.User) error {
	var existingUser model.User
	if err := r.DB.First(&existingUser, userID).Error; err != nil {
		return err
	}

	// Update fields as needed
	existingUser.Name = updatedUser.Name
	existingUser.Phone = updatedUser.Phone
	existingUser.Address = updatedUser.Address // Assuming you want to update the address as well

	return r.DB.Save(&existingUser).Error
}

func (r *UserRepository) DeleteUser(userID uint) error {
	return r.DB.Delete(&model.User{}, userID).Error
}
func (r *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := r.DB.Preload("Address").Find(&users).Error
	return users, err
}

func (r *UserRepository) UserExistsById(id uint) bool {
	var user model.User
	result := r.DB.First(&user, id)
	return result.Error == nil
}
