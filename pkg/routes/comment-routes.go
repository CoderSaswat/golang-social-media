package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"social-media/pkg/handler"
	_ "social-media/pkg/handler"
)

func SetupCommentsRoutes(router *mux.Router) {

	commentHandler := handler.NewCommentHandler()
	router.HandleFunc("/comments", commentHandler.CreateComment).Methods("POST")
	router.HandleFunc("/comments/{id}", commentHandler.GetCommentByID).Methods("GET")
	router.HandleFunc("/comments", commentHandler.GetAllComments).Methods("GET")
	router.HandleFunc("/comments/{id}", commentHandler.UpdateComment).Methods("PUT")
	router.HandleFunc("/comments/posts/{postId}", commentHandler.GetCommentsByPost).Methods("GET")
	router.HandleFunc("/comments/{id}", commentHandler.DeleteComment).Methods("DELETE")

	fmt.Println("comments routes set up successfully")
}
