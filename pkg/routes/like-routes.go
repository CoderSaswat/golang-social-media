package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"social-media/pkg/handler"
)

func SetupLikeRoutes(router *mux.Router) {

	likeHandler := handler.LikeHandler{}
	// Define routes
	router.HandleFunc("/likes", likeHandler.CreateLike).Methods("POST")
	router.HandleFunc("/likes/{id}", likeHandler.GetLikeByID).Methods("GET")
	router.HandleFunc("/likes/{id}", likeHandler.UpdateLike).Methods("PUT")
	router.HandleFunc("/likes/{id}", likeHandler.DeleteLike).Methods("DELETE")
	router.HandleFunc("/likes", likeHandler.GetAllLikes).Methods("GET")
	router.HandleFunc("/likes/posts/{postId}", likeHandler.GetLikesByPostID).Methods("GET")
	router.HandleFunc("/likes/posts/{postId}/like-count", likeHandler.GetLikeCountByPostID).Methods("GET")

	fmt.Println("like routes set up successfully")
}
