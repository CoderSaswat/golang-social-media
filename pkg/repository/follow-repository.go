package repository

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
)

type FollowRepository struct {
	db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{db: db}
}
func (r FollowRepository) IsFollowing(userID uint, targetUserId uint) bool {
	var follow model.Follow
	// Check if a record exists with the given userID and targetUserId
	result := r.db.Where("user_id = ? AND target_user_id = ?", userID, targetUserId).First(&follow)

	// If no error and a record is found, the user is following
	return result.Error == nil && result.RowsAffected > 0
}

func (r FollowRepository) Follow(follow *model.Follow) error {
	return r.db.Save(&follow).Error
}

func (r FollowRepository) GetByUserIdAndTargetUserId(userID uint, targetUserId uint) model.Follow {
	var follow model.Follow
	// Check if a record exists with the given userID and targetUserId
	r.db.Where("user_id = ? AND target_user_id = ?", userID, targetUserId).First(&follow)

	// If no error and a record is found, the user is following
	return follow
}

func (r FollowRepository) DeleteFollowById(followID uint) error {
	var follow model.Follow

	// Find the follow record by ID
	result := r.db.First(&follow, followID)
	if result.Error != nil {
		// Handle the error (record not found, or other database error)
		return result.Error
	}

	// Delete the follow record
	result = r.db.Delete(&follow)
	if result.Error != nil {
		// Handle the error (failed to delete the record)
		return result.Error
	}

	return nil
}

func (r FollowRepository) GetFollowByUserId(userId uint) []model.Follow {
	var follows []model.Follow
	// Find all Follow records where UserID matches
	r.db.Preload("TargetUser").Where("user_id = ?", userId).Find(&follows)
	return follows
}

func (r FollowRepository) GetFollowByTargetUserId(userId uint) []model.Follow {
	var follows []model.Follow
	// Find all Follow records where UserID matches
	r.db.Preload("User").Where("target_user_id = ?", userId).Find(&follows)
	return follows
}
