package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"social-media/pkg/config"
	"social-media/pkg/model"
	"social-media/pkg/service"
	"social-media/pkg/service-impl"
	"social-media/pkg/utils"
	"strconv"
)

type NotificationHandler struct {
}

var notificationService service.NotificationService

func init() {
	db := config.GetDb()
	notificationService = service_impl.NewNotificationService(db)
}

func (h *NotificationHandler) CreateNotification(w http.ResponseWriter, r *http.Request) {
	var notification model.Notification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&notification); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := notificationService.CreateNotification(&notification); err != nil {
		http.Error(w, "Error creating notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *NotificationHandler) GetNotificationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	notification, err := notificationService.GetNotificationByID(uint(notificationID))
	if err != nil {
		http.Error(w, "Notification not found", http.StatusNotFound)
		return
	}

	utils.JsonResponse(w, notification)
}

func (h *NotificationHandler) GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	notifications, err := notificationService.GetAllNotifications()
	if err != nil {
		http.Error(w, "Error fetching notifications", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, notifications)
}

func (h *NotificationHandler) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notificationID, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid notification ID", http.StatusBadRequest)
		return
	}

	if err := notificationService.DeleteNotification(uint(notificationID)); err != nil {
		http.Error(w, "Error deleting notification", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *NotificationHandler) GetAllNotificationsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.ParseUint(vars["userId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	notifications, err := notificationService.GetAllNotificationsByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Error fetching notifications", http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, notifications)
}
