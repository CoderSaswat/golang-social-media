// routes/user-routes.go
package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"social-media/pkg/handler"
	_ "social-media/pkg/handler"
)

func SetupUserRoutes(r *mux.Router) {

	userHandler := handler.UserHandler{}
	// User CRUD routes
	r.HandleFunc("/users", userHandler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users", userHandler.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/follow-or-unfollow", userHandler.FollowOrUnfollwHandler).Methods("POST")
	r.HandleFunc("/users/{id}/follow-or-unfollow", userHandler.GetFollowingsAndFollowersInfo).Methods("GET")

	fmt.Println("user routes set up successfully")
}
