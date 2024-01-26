package service_impl

import (
	"gorm.io/gorm"
	"social-media/pkg/model"
	"social-media/pkg/repository"
)

type NotificationServiceImpl struct {
	notificationRepository *repository.NotificationRepository
}

func NewNotificationService(db *gorm.DB) *NotificationServiceImpl {
	notificationRepository := repository.NewNotificationRepository(db)
	return &NotificationServiceImpl{notificationRepository: notificationRepository}
}

func (s *NotificationServiceImpl) CreateNotification(notification *model.Notification) error {
	return s.notificationRepository.CreateNotification(notification)
}

func (s *NotificationServiceImpl) GetNotificationByID(notificationID uint) (*model.Notification, error) {
	return s.notificationRepository.GetNotificationByID(notificationID)
}

func (s *NotificationServiceImpl) GetAllNotifications() ([]model.Notification, error) {
	return s.notificationRepository.GetAllNotifications()
}

func (s *NotificationServiceImpl) DeleteNotification(notificationID uint) error {
	return s.notificationRepository.DeleteNotification(notificationID)
}

func (s *NotificationServiceImpl) GetAllNotificationsByUserID(userID uint) ([]model.Notification, error) {
	return s.notificationRepository.GetAllNotificationsByUserID(userID)
}
