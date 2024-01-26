package repository

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
)

type NotificationRepository struct {
	DB *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (r *NotificationRepository) CreateNotification(notification *model.Notification) error {
	return r.DB.Create(notification).Error
}

func (r *NotificationRepository) GetNotificationByID(notificationID uint) (*model.Notification, error) {
	var notification model.Notification
	err := r.DB.First(&notification, notificationID).Error
	return &notification, err
}

func (r *NotificationRepository) GetAllNotifications() ([]model.Notification, error) {
	var notifications []model.Notification
	err := r.DB.Find(&notifications).Error
	return notifications, err
}

func (r *NotificationRepository) DeleteNotification(notificationID uint) error {
	return r.DB.Delete(&model.Notification{}, notificationID).Error
}

func (r *NotificationRepository) GetAllNotificationsByUserID(userID uint) ([]model.Notification, error) {
	var notifications []model.Notification
	err := r.DB.Where("user_id = ?", userID).Find(&notifications).Error
	return notifications, err
}
