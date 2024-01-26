package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"social-media/pkg/handler"
	_ "social-media/pkg/handler"
)

func SetupPostRoutes(router *mux.Router) {

	postHandler := handler.PostHandler{}
	// User CRUD routes

	router.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", postHandler.GetPostByID).Methods("GET")
	router.HandleFunc("/posts", postHandler.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", postHandler.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", postHandler.DeletePost).Methods("DELETE")
	router.HandleFunc("/posts/users/{userId}", postHandler.GetAllPostsByUserID).Methods("GET")

	fmt.Println("post routes set up successfully")
}
