package main

import (
	"fmt"
	"log"
	"net/http"
	"social-media/pkg/config"
	"social-media/pkg/routes"
)

func main() {
	//db := config.GetDb()
	//err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}, &model.Image{}, &model.Like{}, &model.Notification{}, &model.Address{}, &model.Follow{})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err := config.SqlServerClient.Ping()
	//if err != nil {
	//	return
	//}
	cli := config.GetMsSqlClient()
	err := cli.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cli.Ping())
	fmt.Println("Connected to SQL Server!")

	fmt.Println("database migrated successfully")
	router := routes.InitializeRoutes()
	err = http.ListenAndServe(":8080", router)
	if err == nil {
		fmt.Println("Server started on port 8080")
	} else {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
