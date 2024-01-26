package main

import (
	"fmt"
	"net/http"
	"social-media/pkg/config"
	"social-media/pkg/model"
	"social-media/pkg/routes"
)

func main() {
	db := config.GetDb()
	err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}, &model.Image{}, &model.Like{}, &model.Notification{}, &model.Address{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("database migrated successfully")
	router := routes.InitializeRoutes()
	err = http.ListenAndServe(":8080", router)
	if err == nil {
		fmt.Println("Server started on port 8080")
	} else {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
