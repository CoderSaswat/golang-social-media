package service

import "social-media/pkg/model"

type NotificationService interface {
	CreateNotification(notification *model.Notification) error
	GetNotificationByID(notificationID uint) (*model.Notification, error)
	GetAllNotifications() ([]model.Notification, error)
	DeleteNotification(notificationID uint) error
	GetAllNotificationsByUserID(userID uint) ([]model.Notification, error)
}
