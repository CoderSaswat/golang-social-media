package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"social-media/pkg/handler"
)

func SetupRedisRoutes(r *mux.Router) {

	redisHandler := handler.RedisHandler{}
	r.HandleFunc("/redis", redisHandler.PostData).Methods("POST")
	fmt.Println("redis routes set up successfully")
}
