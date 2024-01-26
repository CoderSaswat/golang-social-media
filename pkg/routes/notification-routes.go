package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"social-media/pkg/handler"
)

func SetupNotificationRoutes(router *mux.Router) {

	notificationHandler := handler.NotificationHandler{}
	// User CRUD routes

	router.HandleFunc("/notifications", notificationHandler.CreateNotification).Methods("POST")
	router.HandleFunc("/notifications/{id}", notificationHandler.GetNotificationByID).Methods("GET")
	router.HandleFunc("/notifications", notificationHandler.GetAllNotifications).Methods("GET")
	router.HandleFunc("/notifications/{id}", notificationHandler.DeleteNotification).Methods("DELETE")
	router.HandleFunc("/notifications/users/{userId}", notificationHandler.GetAllNotificationsByUserID).Methods("GET")

	fmt.Println("notification routes set up successfully")
}
